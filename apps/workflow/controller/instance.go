package controller

import (

	"github.com/gin-gonic/gin"
	"go-web-mini/apps/workflow/service"
	"go-web-mini/apps/workflow/vo"
	"net/http"
)


//api.POST("/instance", CreateInstance)
//api.GET("/instance/:instance_id", GetInstance)
//api.GET("/instance/:instance_id/variables", GetInstanceVariables)
//api.POST("/instance/:instance_id/variables", SetInstanceVariables)
//api.GET("/instance/:instance_id/commands", GetAvailableCommands)
//api.POST("/instance/available/commands/batch", GetInsListAvailableCommands)
//api.POST("/instance/:instance_id/command", CommandExecution)
//api.POST("/instance/command/optimize/run", CommandExecutionOptimize)
//api.POST("/instance/command/batch/run", CommandExecutionBatch)
//api.GET("/instance/:instance_id/transitions", GetInstanceTransitions)
//api.POST("/instance/command", GetInstanceCommand)
//api.POST("/instance/tasks", GetInstanceTasks)
//api.POST("/instance/schemacode", GetInstanceSchemaCode)
//api.POST("/instance/schemacode/set", SetInstanceSchemaCode)

//引用上面的注释，可以看到这里的路由和上面的注释是一一对应的，这里的路由就是对应的上面的注释的路由

//@description 创建一个流程实例
//@router /apibpmn
type InstanceController struct{

}

//@description 创建一个流程实例
//@router /instance [post]
func (cc *InstanceController)CreateInstance(c *gin.Context) {
	ctx := c.Request.Context()
	var createReq vo.CreateInstance
	if err := easyBind(c, &createReq); err != nil {
		return
	}
	instService := service.NewInstanceService()
	instanceId, e := instService.CreateInstance(ctx, createReq.SchemeCode, service.VarJsonToMap(createReq.Variables))
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, vo.CreateInstanceResp{InstanceID: instanceId}))
}

//@description 获取一个流程实例
//@router /instance/:instance_id [get]
func (cc *InstanceController)GetInstance(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.GetInstanceReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	instService := service.NewInstanceService()
	instanceInfo, definitions, e := instService.GetInstance(ctx, req.InstanceId)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	serviceTasks := convertServiceTasks(definitions.Process.ServiceTasks)
	commands := convertCommands(definitions.Process.IntermediateCatchEvent)
	variables := convertVariables(instanceInfo.GetVariables())
	resp := vo.GetInstanceResp{
		InstanceID:    instanceInfo.Id,
		SchemeCode:    instanceInfo.SchemeCode,
		State:         instanceInfo.State,
		CurrentTaskID: instanceInfo.CurrentTaskId,
		ServiceTasks:  serviceTasks,
		Commands:      commands,
		Variables:     variables,
	}
	for _, serviceTask := range serviceTasks {
		if serviceTask.TaskId == instanceInfo.CurrentTaskId {
			resp.CurrentTaskName = serviceTask.TaskName
			break
		}
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, resp))
}

//@description 获取一个流程实例的变量
//@router /instance/schemacode [post]
func (cc *InstanceController)GetInstanceSchemaCode(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.InstanceSchemaCodeReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	instService := service.NewInstanceService()
	insIdMapSchemaCode, e := instService.GetInstanceSchemaCode(ctx, req.InstanceIdList)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	var instanceSchemaCodeList []vo.InstanceSchemaCode
	for key, val := range insIdMapSchemaCode {
		instanceSchemaCodeList = append(instanceSchemaCodeList, vo.InstanceSchemaCode{
			InstanceID:        key,
			SchemaCode:        val.SchemaCode,
			InstanceCurTaskId: val.InstanceCurTaskId,
		})
	}
	resp := vo.GetInstanceSchemaCodeResp{
		InstanceSchemaCodeList: instanceSchemaCodeList,
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, resp))
}

//@description 设置一个流程实例的变量
//@router /instance/schemacode/set [post]
func (cc *InstanceController)SetInstanceSchemaCode(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.InstanceSchemaCodeReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	instService := service.NewInstanceService()
	e := instService.SetInstanceSchemaCode(ctx, req.InstanceIdList, req.SchemaCode)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, nil))
}

//@description 获取一个实例的任务列表
//@router /instance/tasks [post]
func (cc *InstanceController)GetInstanceTasks(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.GetInstanceTasksReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	instService := service.NewInstanceService()
	instanceTabMap, definitionsMap, e := instService.GetInstanceTasks(ctx, req.InstanceIdList)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	var instanceTasksList []vo.InstanceTasks
	for key, instanceTab := range instanceTabMap {
		definition, ok := definitionsMap[key]
		if !ok {
			continue
		}
		currentTaskId := instanceTab.CurrentTaskId
		instanceTasks := vo.InstanceTasks{
			InstanceID: instanceTab.Id,
		}
		tasks := convertServiceTasks(definition.Process.ServiceTasks)
		for _, task := range tasks {
			if task.TaskId == currentTaskId {
				instanceTasks.CurrentTaskName = task.TaskName
				break
			}
		}
		instanceTasks.Tasks = tasks
		instanceTasksList = append(instanceTasksList, instanceTasks)
	}

	resp := vo.GetInstanceTasksResp{
		InstanceTasksList: instanceTasksList,
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, resp))
}

//@description 获取实例命令
//@router /instance/command [post]
func (cc *InstanceController)GetInstanceCommand(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.GetInstanceCommandReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	instService := service.NewInstanceService()
	idMapCatchEvent, e := instService.GetInstanceCommand(ctx, req.InstanceIdList, req.CommandKey)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	var instanceCommandList []vo.InstanceCommand
	for id, catchEvent := range idMapCatchEvent {
		extendProperties := catchEvent.ExtensionElements.Properties.Properties
		var variables []vo.Variables
		for _, property := range extendProperties {
			variables = append(variables, vo.Variables{
				Name:  property.Name,
				Value: property.Value,
			})
		}
		command := vo.Command{Id: catchEvent.Id, Key: catchEvent.Name, Variables: variables}
		instanceCommandList = append(instanceCommandList, vo.InstanceCommand{
			InstanceID: id,
			Command:    command,
		})
	}
	resp := vo.GetInstanceCommandResp{
		InstanceCommandList: instanceCommandList,
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, resp))
}

//@description 获取实例流转记录
//@router /instance/:instance_id/transitions [post]
func (cc *InstanceController)GetInstanceTransitions(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.GetInstanceTransitions
	if err := easyBind(c, &req); err != nil {
		return
	}
	instService := service.NewInstanceService()
	transitions, e := instService.GetTransitions(ctx, req.InstanceId)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	ts := make([]vo.InstanceTransition, len(transitions))
	for i, transition := range transitions {
		ts[i] = vo.InstanceTransition{
			InstanceId:     transition.InstanceId,
			IdentityId:     transition.IdentityId,
			FromStateName:  transition.FromActivityName,
			ToStateName:    transition.ToActivityName,
			CommandName:    transition.CommandName,
			Variables:      convertVariables(transition.GetVariables()),
			TransitionTime: transition.Ctime,
		}
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, vo.InstanceTransitions{Transitions: ts}))
}
