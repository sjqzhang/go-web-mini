package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-web-mini/apps/workflow/bpmn_engine"
	"go-web-mini/apps/workflow/errcode"
	"go-web-mini/apps/workflow/listener"
	"go-web-mini/apps/workflow/model"
	"go-web-mini/apps/workflow/repository"
	"go-web-mini/apps/workflow/spec/BPMN20"
	"go-web-mini/apps/workflow/spec/BPMN20/process_instance"
	"go-web-mini/global"
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
	schemeTab, err := schemeService.Get(ctx, schemeCode)
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
	//TODO 事务
	//err = global.DB.Transaction().Execute(ctx, func(c context.Context) error {
	//	err := i.InstanceRepo.New(c, &instanceTab)
	//	if err != nil {
	//		return err
	//	}
	//	err = i.saveTransitions(c, pi)
	//	if e != nil {
	//		return err
	//	}
	//	return nil
	//})
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
	schemeTab, err := schemeService.Get(ctx, instanceTab.SchemeCode)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"scheme_code": instanceTab.SchemeCode,
		}).Errorf("get scheme from DB err: %+v", err)
		return nil, nil, errcode.Wrap(errcode.DBError, err)
	}
	engine := bpmn_engine.New("")
	processInfo, _ := engine.LoadFromBytes([]byte(schemeTab.Scheme))

	processDefinition := processInfo.Definitions()
	return instanceTab, &processDefinition, nil
}

func (i *InstanceService) GetInstanceEngine(ctx context.Context, instance *model.InstanceTab) (*bpmn_engine.BpmnEngineState, error) {
	schemeService := NewSchemeService()
	schemeTab, err := schemeService.Get(ctx, instance.SchemeCode)
	if err != nil {
		return nil, err
	}
	engine := bpmn_engine.New("")
	// add listeners
	engine.AddListener("orApprovalHandler", listener.OrApprovalListener{})
	engine.AddListener("conditionsExpression", listener.ConditionsExpressionChecker{})
	// unmarshal processInfo
	processInfo, err := engine.LoadFromBytes([]byte(schemeTab.Scheme))
	if err != nil {
		return nil, err
	}
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

func (i *InstanceService) CommandExecute(ctx context.Context, instanceId int, command, identityId string, variables map[string]interface{}) (e errcode.Exception) {
	instance, err := i.InstanceRepo.Get(ctx, instanceId)
	if err != nil {
		return errcode.Wrap(errcode.DBError, err)
	}
	if instance.State == string(process_instance.COMPLETED) {
		return nil
	}

	engine, err := i.GetInstanceEngine(ctx, instance)
	if err != nil {
		return errcode.Wrap(errcode.EngineError, err)
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
		return errcode.CommandIsAvailable
	}

	// Lock
	lockerService := NewLockerService()
	locker, err := lockerService.Lock(ctx, strconv.Itoa(instanceId), 10*time.Second)
	if err != nil {
		return errcode.Wrap(errcode.LockFail, err)
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
		return errcode.Wrap(errcode.EngineError, err)
	}
	_, err = engine.RunOrContinueInstance(pi.GetInstanceKey())
	if err != nil {
		return errcode.Wrap(errcode.EngineError, err)
	}
	instance.State = string(pi.GetState())
	instance.CurrentTaskId = pi.GetCurrentTask().GetId()

	// Persistence
	//TODO: 事务
	err = global.GetTransaction().Execute(ctx, func(c context.Context) error {
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
		err = i.saveTransitions(c, pi)
		return err
	})
	if err != nil {
		return errcode.Wrap(errcode.DBError, err)
	}

	return nil
}

func (i *InstanceService) GetTransitions(ctx context.Context, instanceId int) (transitions []model.TransitionsTab, e errcode.Exception) {
	transitions, err := i.TransitionsRepo.QueryByInstanceId(ctx, instanceId)
	if err != nil {
		return nil, errcode.Wrap(errcode.DBError, err)
	}
	return transitions, nil
}

func (i *InstanceService) saveTransitions(ctx context.Context, pi *bpmn_engine.ProcessInstanceInfo) (e error) {
	instanceId := int(pi.GetInstanceKey())
	identityId, _ := pi.GetVariable(IdentityId).(string)
	tasks := pi.GetWalkedTasks()
	if len(tasks) <= 1 {
		return nil
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

	err := global.GetTransaction().Execute(ctx, func(c context.Context) error {
		for _, m := range ts {
			err := i.TransitionsRepo.New(c, &m)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
