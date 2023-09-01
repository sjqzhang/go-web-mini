package service

import (
	"context"
	"errors"
	"fmt"
	"go-web-mini/apps/workflow/bpmn_engine"
	"go-web-mini/apps/workflow/errcode"
	"go-web-mini/apps/workflow/listener"
	"go-web-mini/apps/workflow/model"
	"go-web-mini/apps/workflow/repository"
	"go-web-mini/apps/workflow/spec/BPMN20"
	"go-web-mini/apps/workflow/spec/BPMN20/process_instance"
	"go-web-mini/apps/workflow/vo"
	"go-web-mini/global"

	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const IdentityId = "identity_id"

type InstanceService struct {
	SchemeRepo       repository.SchemeRepo
	InstanceRepo     repository.InstanceRepo
	CaughtEventsRepo repository.CaughtEventsRepo
	TransitionsRepo  repository.TransitionsRepo
}

func NewInstanceService() *InstanceService {
	return &InstanceService{
		SchemeRepo:       &repository.SchemeDBRepo{},
		InstanceRepo:     &repository.InstanceDBRepo{},
		CaughtEventsRepo: &repository.CaughtEventsDBRepo{},
		TransitionsRepo:  &repository.TransitionsDBRepo{},
	}
}

func (i *InstanceService) CreateInstance(ctx context.Context, schemeCode string, varMap map[string]interface{}) (instanceId int, e errcode.Exception) {
	schemeService := NewSchemeService()
	schemeTab, _, err := schemeService.Get(ctx, schemeCode)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"scheme_code": schemeCode,
		}).Errorf("get scheme from DB err: %+v", err)
		return 0, errcode.Wrap(errcode.DBError, err)
	}
	bpmnEngine := bpmn_engine.New(schemeCode + "-engine")
	process, _ := bpmnEngine.LoadFromBytes([]byte(schemeTab.Scheme))
	pi, err := bpmnEngine.CreateInstance(process.ProcessKey, varMap)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"scheme_code": schemeCode,
		}).Errorf("create instance err: %+v", err)
		return 0, errcode.Wrap(errcode.EngineError, err)
	}

	_, err = bpmnEngine.RunOrContinueInstance(pi.GetInstanceKey())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"instance_id": pi.GetInstanceKey(),
			"scheme_code": schemeCode,
		}).Errorf("run instance err: %+v", err)
		return 0, errcode.Wrap(errcode.EngineError, err)
	}

	instanceTab := model.InstanceTab{
		Id:            int(pi.GetInstanceKey()),
		SchemeCode:    schemeCode,
		State:         string(pi.GetState()),
		CurrentTaskId: pi.GetCurrentTask().GetId(),
	}
	instanceTab.SetVariables(varMap)
	err = global.GetTransaction().Execute(ctx, func(c context.Context) error {
		err := i.InstanceRepo.New(c, &instanceTab)
		if err != nil {
			return err
		}
		_, err = i.saveTransitions(c, pi)
		if e != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"instance_id": pi.GetInstanceKey(),
			"scheme_code": schemeCode,
		}).Errorf("save instance err: %+v", err)
		return 0, errcode.Wrap(errcode.DBError, err)
	}

	return instanceTab.Id, nil
}

func (i *InstanceService) GetInstance(ctx context.Context, instanceId int) (*model.InstanceTab, *BPMN20.TDefinitions, errcode.Exception) {
	instanceTab, err := i.InstanceRepo.Get(ctx, instanceId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"instance_id": instanceId,
		}).Errorf("get instance from DB err: %+v", err)
		return nil, nil, errcode.Wrap(errcode.DBError, err)
	}
	schemeService := NewSchemeService()
	_, processInfo, err := schemeService.Get(ctx, instanceTab.SchemeCode)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"scheme_code": instanceTab.SchemeCode,
		}).Errorf("get scheme from DB err: %+v", err)
		return nil, nil, errcode.Wrap(errcode.DBError, err)
	}
	processDefinition := processInfo.Definitions()
	return instanceTab, &processDefinition, nil
}

func (i *InstanceService) GetInstanceEngine(ctx context.Context, instance *model.InstanceTab) (*bpmn_engine.BpmnEngineState, error) {
	schemeService := NewSchemeService()
	_, processInfo, err := schemeService.Get(ctx, instance.SchemeCode)
	if err != nil {
		return nil, err
	}
	engine := bpmn_engine.New("")
	// add listeners
	engine.AddListener("orApprovalHandler", listener.OrApprovalListener{})
	engine.AddListener("conditionsExpression", listener.ConditionsExpressionChecker{})
	var currentTask BPMN20.TServiceTask
	for _, serviceTask := range processInfo.Definitions().Process.ServiceTasks {
		if serviceTask.Id == instance.CurrentTaskId {
			currentTask = serviceTask
			break
		}
	}

	// unmarshal processInstanceInfo
	events, err := i.CaughtEventsRepo.QueryByInstanceId(ctx, instance.Id)
	if err != nil {
		return nil, err
	}
	catchEvents := make([]bpmn_engine.CatchEvent, len(events))
	for i, event := range events {
		catchEvents[i] = bpmn_engine.CatchEvent{
			Name:       event.EventName,
			CaughtAt:   time.Unix(int64(event.Ctime), 0),
			IsConsumed: event.HasConsumed(),
		}
	}
	processInstance := bpmn_engine.NewProcessInstanceInfo(
		processInfo, int64(instance.Id), instance.GetVariables(),
		time.Unix(int64(instance.Ctime), 0),
		process_instance.State(instance.State), catchEvents,
		[]BPMN20.BaseElement{currentTask}, currentTask,
	)
	engine.SetProcessInstance(processInstance)
	return &engine, nil
}

func (i *InstanceService) GetInstanceSchemaCode(ctx context.Context, instanceIds []int) (map[int]*vo.InstanceSchemaCode, errcode.Exception) {
	instanceTabList, err := i.InstanceRepo.GetList(ctx, instanceIds)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"instance_ids": instanceIds,
		}).Errorf("get instance from DB err: %+v", err)
		return nil, errcode.Wrap(errcode.DBError, err)
	}
	insIdMapSchemaCode := make(map[int]*vo.InstanceSchemaCode)
	for _, instanceTab := range instanceTabList {
		//insIdMapSchemaCode[instanceTab.Id] = instanceTab.SchemeCode
		insIdMapSchemaCode[instanceTab.Id] = &vo.InstanceSchemaCode{
			InstanceID:        instanceTab.Id,
			SchemaCode:        instanceTab.SchemeCode,
			InstanceCurTaskId: instanceTab.CurrentTaskId,
		}
	}
	return insIdMapSchemaCode, nil
}

func (i *InstanceService) SetInstanceSchemaCode(ctx context.Context, instanceIds []int, schemaCode string) errcode.Exception {
	err := i.InstanceRepo.UpdateSchemaCode(ctx, instanceIds, schemaCode)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"instance_ids": instanceIds,
		}).Errorf("SetInstanceSchemaCode from DB err: %+v", err)
		return errcode.Wrap(errcode.DBError, err)
	}
	return nil
}

func (i *InstanceService) GetInstanceTasks(ctx context.Context, instanceIds []int) (map[int]*model.InstanceTab, map[int]*BPMN20.TDefinitions, errcode.Exception) {
	instanceTabList, err := i.InstanceRepo.GetList(ctx, instanceIds)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"instance_id": instanceIds,
		}).Errorf("get instance from DB err: %+v", err)
		return nil, nil, errcode.Wrap(errcode.DBError, err)
	}
	instanceIdMapTab := make(map[int]*model.InstanceTab)
	instanceIdMapDefinition := make(map[int]*BPMN20.TDefinitions)
	schemeService := NewSchemeService()
	for _, instanceTab := range instanceTabList {
		_, processInfo, err := schemeService.Get(ctx, instanceTab.SchemeCode)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"scheme_code": instanceTab.SchemeCode,
			}).Errorf("get scheme from DB err: %+v", err)
			return nil, nil, errcode.Wrap(errcode.DBError, err)
		}
		processDefinition := processInfo.Definitions()
		instanceIdMapTab[instanceTab.Id] = instanceTab
		instanceIdMapDefinition[instanceTab.Id] = &processDefinition
	}

	return instanceIdMapTab, instanceIdMapDefinition, nil
}

func (i *InstanceService) GetInstanceCommand(ctx context.Context, instanceIds []int, commandKey string) (idMapCatchEvent map[int]*BPMN20.TIntermediateCatchEvent, e errcode.Exception) {
	// 获取instanceId对应的schemaCode
	instanceTabList, err := i.InstanceRepo.GetList(ctx, instanceIds)
	if err != nil {
		return nil, errcode.Wrap(errcode.DBError, err)
	}
	idMapCatchEvent = make(map[int]*BPMN20.TIntermediateCatchEvent)
	for _, instanceTab := range instanceTabList {
		_, processInfo, err := NewSchemeService().Get(ctx, instanceTab.SchemeCode)
		if err != nil {
			return nil, errcode.Wrap(errcode.DBError, err)
		}
		var catchEvent *BPMN20.TIntermediateCatchEvent
		for _, item := range processInfo.Definitions().Process.IntermediateCatchEvent {
			if item.Name == commandKey {
				catchEvent = &item
				break
			}
		}
		if catchEvent == nil {
			return nil, errcode.Wrap(errcode.ServiceError,
				fmt.Errorf("schema code = %s no commandKey = %s", instanceTab.SchemeCode, commandKey))
		}
		idMapCatchEvent[instanceTab.Id] = catchEvent
	}
	return idMapCatchEvent, nil
}

func (i *InstanceService) GetAvailableCommands(ctx context.Context, instanceId int, identityId string) (events []*BPMN20.TIntermediateCatchEvent, e errcode.Exception) {
	instance, err := i.InstanceRepo.Get(ctx, instanceId)
	if err != nil {
		return nil, errcode.Wrap(errcode.DBError, err)
	}
	if instance.State == string(process_instance.COMPLETED) {
		return nil, nil
	}
	engine, err := i.GetInstanceEngine(ctx, instance)
	if err != nil {
		return nil, errcode.Wrap(errcode.EngineError, err)
	}
	pi := engine.FindProcessInstanceById(int64(instanceId))
	pi.SetVariable(IdentityId, identityId)
	return engine.GetNextAvailableCommands(pi), nil
}

//@description: 获取多个实例的可用命令
//@router /instance/available/commands/batch [post]
func (i *InstanceService) GetInsListAvailableCommands(ctx context.Context, instanceIds []int, identityId string) (insIdMapEvents map[int][]*BPMN20.TIntermediateCatchEvent, e errcode.Exception) {
	instanceList, err := i.InstanceRepo.GetList(ctx, instanceIds)
	if err != nil {
		return nil, errcode.Wrap(errcode.DBError, err)
	}
	insIdMapEvents = make(map[int][]*BPMN20.TIntermediateCatchEvent)
	for _, instance := range instanceList {
		if instance.State == string(process_instance.COMPLETED) {
			return nil, nil
		}
		engine, err := i.GetInstanceEngine(ctx, instance)
		if err != nil {
			return nil, errcode.Wrap(errcode.EngineError, err)
		}
		pi := engine.FindProcessInstanceById(int64(instance.Id))
		pi.SetVariable(IdentityId, identityId)
		events := engine.GetNextAvailableCommands(pi)
		insIdMapEvents[instance.Id] = events
	}
	return insIdMapEvents, nil
}

//func (i *InstanceService) CommandExecute(ctx context.Context, instanceId int, command, identityId string, variables map[string]interface{}) (transitionsId []int, e errcode.Exception) {
//	instance, err := i.InstanceRepo.Get(ctx, instanceId)
//	if err != nil {
//		return nil, errcode.Wrap(errcode.DBError, err)
//	}
//	if instance.State == string(process_instance.COMPLETED) {
//		return []int{}, nil
//	}
//
//	engine, err := i.GetInstanceEngine(ctx, instance)
//	if err != nil {
//		return nil, errcode.Wrap(errcode.EngineError, err)
//	}
//	pi := engine.FindProcessInstanceById(int64(instanceId))
//	if identityId != "" {
//		pi.SetVariable(IdentityId, identityId)
//	}
//	// Check target command and variables are available
//	events := engine.GetNextAvailableCommands(pi)
//	commandIsAvailable, variablesIsAvailable := false, true
//	for _, event := range events {
//		if event.Name != command {
//			continue
//		}
//		commandIsAvailable = true
//		properties := event.GetExtensionProperties()
//
//		for _, property := range properties {
//			if _, exist := variables[property]; !exist {
//				variablesIsAvailable = false
//				break
//			}
//		}
//		break
//	}
//	if !commandIsAvailable || !variablesIsAvailable {
//		return nil, errcode.CommandIsAvailable
//	}
//
//	// Lock
//	lockerService := NewLockerService()
//	locker, err := lockerService.Lock(ctx, strconv.Itoa(instanceId), 10*time.Second)
//	if err != nil {
//		return nil, errcode.Wrap(errcode.LockFail, err)
//	}
//	defer func() {
//		if locker == nil {
//			return
//		}
//		err := lockerService.Unlock(ctx, locker.Id)
//		if err != nil {
//			logrus.WithFields(logrus.Fields{
//				"instance_id": instanceId,
//			}).Errorf("unlock fail: %+v", err)
//		}
//	}()
//
//	// Set variables and exec command
//	for k, v := range variables {
//		pi.SetVariable(k, v.(string))
//	}
//	err = engine.PublishEventForInstance(pi.GetInstanceKey(), command)
//	if err != nil {
//		return nil, errcode.Wrap(errcode.EngineError, err)
//	}
//	_, err = engine.RunOrContinueInstance(pi.GetInstanceKey())
//	if err != nil {
//		return nil, errcode.Wrap(errcode.EngineError, err)
//	}
//	instance.State = string(pi.GetState())
//	instance.CurrentTaskId = pi.GetCurrentTask().GetId()
//
//	// Persistence
//	err = global.GetTransaction().Execute(ctx, func(c context.Context) error {
//		err := i.InstanceRepo.Update(c, instance)
//		if err != nil {
//			return err
//		}
//		err = i.CaughtEventsRepo.New(c, &model.CaughtEventsTab{
//			InstanceId: instanceId,
//			EventName:  command,
//			IsConsumed: 1,
//		})
//		if err != nil {
//			return err
//		}
//		transitionsId, err = i.saveTransitions(c, pi)
//		return err
//	})
//	if err != nil {
//		return nil, errcode.Wrap(errcode.DBError, err)
//	}
//
//	return transitionsId, nil
//}

func (i *InstanceService) CommandExecute(ctx context.Context, instanceId int, command, identityId string, variables map[string]interface{}) (transitionsId []int, e errcode.Exception) {
	commandEngine, e := i.runCommandByEngine(ctx, instanceId, command, identityId, variables)
	if e != nil {
		return nil, e
	}
	instance := commandEngine.InstanceTab
	pi := commandEngine.ProcessInstanceInfo
	if e != nil {
		return nil, e
	}

	// Persistence
	err := global.GetTransaction().Execute(ctx, func(c context.Context) error {
		err := i.InstanceRepo.Update(c, instance)
		if err != nil {
			return err
		}
		err = i.CaughtEventsRepo.New(c, &model.CaughtEventsTab{
			InstanceId: instanceId,
			EventName:  command,
			IsConsumed: 1,
		})
		if err != nil {
			return err
		}
		transitionsId, err = i.saveTransitions(c, pi)
		return err
	})
	if err != nil {
		return nil, errcode.Wrap(errcode.DBError, err)
	}

	return transitionsId, nil
}

type CommandEngine struct {
	InstanceTab         *model.InstanceTab
	ProcessInstanceInfo *bpmn_engine.ProcessInstanceInfo
	Command             string
}

func (i *InstanceService) runCommandByEngine(ctx context.Context, instanceId int, command, identityId string, variables map[string]interface{}) (*CommandEngine, errcode.Exception) {
	instance, err := i.InstanceRepo.Get(ctx, instanceId)
	if err != nil {
		return nil, errcode.Wrap(errcode.DBError, err)
	}
	if instance.State == string(process_instance.COMPLETED) {
		return nil, errcode.Wrap(errcode.CommandIsAvailable, errors.New("already Completed"))
	}

	engine, err := i.GetInstanceEngine(ctx, instance)
	if err != nil {
		return nil, errcode.Wrap(errcode.EngineError, err)
	}
	pi := engine.FindProcessInstanceById(int64(instanceId))
	if identityId != "" {
		pi.SetVariable(IdentityId, identityId)
	}
	// Check target command and variables are available
	events := engine.GetNextAvailableCommands(pi)
	commandIsAvailable, variablesIsAvailable := false, true
	for _, event := range events {
		if event.Name != command {
			continue
		}
		commandIsAvailable = true
		properties := event.GetExtensionProperties()

		for _, property := range properties {
			if _, exist := variables[property]; !exist {
				variablesIsAvailable = false
				break
			}
		}
		break
	}
	if !commandIsAvailable || !variablesIsAvailable {
		return nil, errcode.CommandIsAvailable
	}

	// Lock
	lockerService := NewLockerService()
	locker, err := lockerService.Lock(ctx, strconv.Itoa(instanceId), 10*time.Second)
	if err != nil {
		return nil, errcode.Wrap(errcode.LockFail, err)
	}
	defer func() {
		if locker == nil {
			return
		}
		err := lockerService.Unlock(ctx, locker.Id)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"instance_id": instanceId,
			}).Errorf("unlock fail: %+v", err)
		}
	}()

	// Set variables and exec command
	for k, v := range variables {
		pi.SetVariable(k, v.(string))
	}
	err = engine.PublishEventForInstance(pi.GetInstanceKey(), command)
	if err != nil {
		return nil, errcode.Wrap(errcode.EngineError, err)
	}
	_, err = engine.RunOrContinueInstance(pi.GetInstanceKey())
	if err != nil {
		return nil, errcode.Wrap(errcode.EngineError, err)
	}
	instance.State = string(pi.GetState())
	instance.CurrentTaskId = pi.GetCurrentTask().GetId()
	return &CommandEngine{
		InstanceTab:         instance,
		ProcessInstanceInfo: pi,
		Command:             command,
	}, nil
}

func (i *InstanceService) CommandExecuteBatch(ctx context.Context, reqBatch vo.CommandExecutionBatchReq) (insIdMapTransitionsIds map[int][]int, e errcode.Exception) {
	commandEngineList := make([]*CommandEngine, 0, len(reqBatch.CommandExecutionBatchList))
	for _, commandExecution := range reqBatch.CommandExecutionBatchList {
		variables := VarJsonToMap(commandExecution.Variables)
		commandEngine, e := i.runCommandByEngine(ctx, commandExecution.InstanceId, commandExecution.Command, commandExecution.IdentityId, variables)
		if e != nil {
			logrus.WithFields(logrus.Fields{
				"instance_id": commandExecution.InstanceId,
			}).Errorf("CommandExecuteBatch fail: %+v", e)
			continue
			//return nil, e
		}
		//commandEngineList[index] = commandEngine
		commandEngineList = append(commandEngineList, commandEngine)
	}

	insIdMapTransitionsIds = make(map[int][]int)
	// Persistence
	for _, commandEngine := range commandEngineList {
		errTx := global.GetTransaction().Execute(ctx, func(c context.Context) error {
			err := i.InstanceRepo.Update(c, commandEngine.InstanceTab)
			if err != nil {
				return err
			}
			err = i.CaughtEventsRepo.New(c, &model.CaughtEventsTab{
				InstanceId: commandEngine.InstanceTab.Id,
				EventName:  commandEngine.Command,
				IsConsumed: 1,
			})
			if err != nil {
				return err
			}
			transitionsId, err := i.saveTransitions(c, commandEngine.ProcessInstanceInfo)
			if err != nil {
				return err
			}
			insIdMapTransitionsIds[commandEngine.InstanceTab.Id] = transitionsId
			return nil
		})
		if errTx != nil {
			logrus.WithFields(logrus.Fields{
				"instance_id": commandEngine.InstanceTab.Id,
			}).Errorf("global.GetTransaction().Execute fail: %+v", e)
		}
	}
	//err := global.GetTransaction().Execute(ctx, func(c context.Context) error {
	//	for _, commandEngine := range commandEngineList {
	//		err := i.InstanceRepo.Update(c, commandEngine.InstanceTab)
	//		if err != nil {
	//			return err
	//		}
	//		err = i.CaughtEventsRepo.New(c, &model.CaughtEventsTab{
	//			InstanceId: commandEngine.InstanceTab.Id,
	//			EventName:  commandEngine.Command,
	//			IsConsumed: 1,
	//		})
	//		if err != nil {
	//			return err
	//		}
	//		transitionsId, err := i.saveTransitions(c, commandEngine.ProcessInstanceInfo)
	//		if err != nil {
	//			return err
	//		}
	//		insIdMapTransitionsIds[commandEngine.InstanceTab.Id] = transitionsId
	//	}
	//	return nil
	//})
	//if err != nil {
	//	return nil, errcode.Wrap(errcode.DBError, err)
	//}

	return insIdMapTransitionsIds, nil
}

func (i *InstanceService) CommandExecuteOptimize(ctx context.Context, reqOptimize vo.CommandExecutionOptimizeReq) (transitionsIdList []int, e errcode.Exception) {
	for _, commandExecution := range reqOptimize.CommandList {
		variables := VarJsonToMap(commandExecution.Variables)
		transitionsId, e := i.CommandExecute(ctx, reqOptimize.InstanceId, commandExecution.Command, reqOptimize.IdentityId, variables)
		if e != nil {
			return nil, e
		}
		transitionsIdList = append(transitionsIdList, transitionsId...)
	}
	return transitionsIdList, nil
}

func (i *InstanceService) GetTransitions(ctx context.Context, instanceId int) (transitions []model.TransitionsTab, e errcode.Exception) {
	transitions, err := i.TransitionsRepo.QueryByInstanceId(ctx, instanceId)
	if err != nil {
		return nil, errcode.Wrap(errcode.DBError, err)
	}
	return transitions, nil
}

func (i *InstanceService) GetTransitionsByIds(ctx context.Context, ids []int) (transitions []model.TransitionsTab, e errcode.Exception) {
	transitions, err := i.TransitionsRepo.QueryByIds(ctx, ids)
	if err != nil {
		return nil, errcode.Wrap(errcode.DBError, err)
	}
	return transitions, nil
}

func (i *InstanceService) saveTransitions(ctx context.Context, pi *bpmn_engine.ProcessInstanceInfo) ([]int, error) {
	instanceId := int(pi.GetInstanceKey())
	identityId, _ := pi.GetVariable(IdentityId).(string)
	tasks := pi.GetWalkedTasks()
	if len(tasks) <= 1 {
		return []int{}, nil
	}

	var ts []model.TransitionsTab
	var preTask, afterTask BPMN20.TServiceTask
	ii := 0
	for ; ii < len(tasks); ii++ {
		if tasks[ii].GetType() == BPMN20.ServiceTask {
			preTask = tasks[ii].(BPMN20.TServiceTask)
			break
		}
	}
	for ; ii < len(tasks)-1; ii++ {
		var command BPMN20.TIntermediateCatchEvent
		for j := ii + 1; j < len(tasks); j++ {
			if tasks[j].GetType() == BPMN20.IntermediateCatchEvent {
				command = tasks[j].(BPMN20.TIntermediateCatchEvent)
			}
			if tasks[j].GetType() != BPMN20.ServiceTask {
				continue
			}
			afterTask = tasks[j].(BPMN20.TServiceTask)
			variables := make(map[string]interface{})
			for _, property := range command.GetExtensionProperties() {
				value := pi.GetVariable(property)
				if value != nil {
					variables[property] = value
				}
			}

			m := model.TransitionsTab{
				InstanceId:       instanceId,
				IdentityId:       identityId,
				FromActivityId:   preTask.GetId(),
				FromActivityName: preTask.GetName(),
				ToActivityId:     afterTask.GetId(),
				ToActivityName:   afterTask.GetName(),
				CommandId:        command.GetId(),
				CommandName:      command.GetName(),
			}
			m.SetVariables(variables)
			ts = append(ts, m)
			preTask = afterTask
			ii = j
			break
		}
	}

	var transitionsTabIds []int
	err := global.GetTransaction().Execute(ctx, func(c context.Context) error {
		for _, m := range ts {
			err := i.TransitionsRepo.New(c, &m)
			if err != nil {
				return err
			}
			transitionsTabIds = append(transitionsTabIds, m.Id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return transitionsTabIds, nil
}
