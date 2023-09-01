package controller

import (

	"github.com/gin-gonic/gin"
	"go-web-mini/apps/workflow/service"
	"go-web-mini/apps/workflow/vo"
	"net/http"
)

//@Description command信息
//@router /apibpmn
type CommandController struct {

}

// @description: 获取实例列表可用的命令
// @router /instance/:instance_id/commands [get]

func (cc *CommandController)GetAvailableCommands(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.GetInstanceAvailableCommandsReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	instanceService := service.NewInstanceService()
	events, e := instanceService.GetAvailableCommands(ctx, req.InstanceId, req.IdentityId)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	commands := make([]vo.Command, len(events))
	for i, event := range events {
		commands[i] = vo.Command{
			Id:  event.Id,
			Key: event.Name,
			//Variables: event.GetExtensionProperties(),
		}
		var variables []vo.Variables
		for _, extendProperty := range event.ExtensionElements.Properties.Properties {
			variables = append(variables, vo.Variables{
				Name:  extendProperty.Name,
				Value: extendProperty.Value,
			})
		}
		commands[i].Variables = variables
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, vo.GetInstanceAvailableCommands{Commands: commands}))
}
//@description: 获取实例列表可用的命令
//@router /instance/available/commands/batch [post]
func (cc *CommandController)GetInsListAvailableCommands(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.GetInsListAvailableCommandsReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	instanceService := service.NewInstanceService()
	insIdMapEvents, e := instanceService.GetInsListAvailableCommands(ctx, req.InstanceIds, req.IdentityId)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	insAvailableCommandsList := make([]vo.InsAvailableCommands, 0, 10)
	for insId, events := range insIdMapEvents {
		commands := make([]vo.Command, len(events))
		for i, event := range events {
			commands[i] = vo.Command{
				Id:  event.Id,
				Key: event.Name,
				//Variables: event.GetExtensionProperties(),
			}
			var variables []vo.Variables
			for _, extendProperty := range event.ExtensionElements.Properties.Properties {
				variables = append(variables, vo.Variables{
					Name:  extendProperty.Name,
					Value: extendProperty.Value,
				})
			}
			commands[i].Variables = variables
		}
		insAvailableCommandsList = append(insAvailableCommandsList, vo.InsAvailableCommands{
			InstanceId:  insId,
			CommandList: commands,
		})
	}

	c.JSON(http.StatusOK, vo.NewResp(nil, vo.GetInsListAvailableCommandsResp{
		InsAvailableCommandsList: insAvailableCommandsList,
	}))
}
//@description 执行命令
//@router /instance/:instance_id/command [post]
func (cc *CommandController) CommandExecution(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.CommandExecution
	if err := easyBind(c, &req); err != nil {
		return
	}

	instanceService := service.NewInstanceService()
	transitionsId, e := instanceService.CommandExecute(ctx, req.InstanceId, req.Command, req.IdentityId, service.VarJsonToMap(req.Variables))
	c.JSON(http.StatusOK, vo.NewResp(e, vo.CommandExecResp{
		TransitionIds: transitionsId,
	}))
}
//@description 批量执行命令
//@router /instance/command/batch/run [post]
func (cc *CommandController)CommandExecutionBatch(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.CommandExecutionBatchReq
	if err := easyBind(c, &req); err != nil {
		return
	}

	instanceService := service.NewInstanceService()
	insIdMapTransitionsIds, e := instanceService.CommandExecuteBatch(ctx, req)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	insIdTransitionIdsList := make([]vo.InsIdTransitionIds, 0, 10)
	for insId, transitionsIds := range insIdMapTransitionsIds {
		insIdTransitionIdsList = append(insIdTransitionIdsList, vo.InsIdTransitionIds{
			InstanceId:    insId,
			TransitionIds: transitionsIds,
		})
	}
	c.JSON(http.StatusOK, vo.NewResp(e, vo.CommandExecBatchResp{
		InsIdTransitionIdsList: insIdTransitionIdsList,
	}))
}
//@description 执行命令优化
//@router /instance/command/optimize/run [post]
func (cc *CommandController)CommandExecutionOptimize(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.CommandExecutionOptimizeReq
	if err := easyBind(c, &req); err != nil {
		return
	}

	instanceService := service.NewInstanceService()
	transitionsIdList, e := instanceService.CommandExecuteOptimize(ctx, req)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	c.JSON(http.StatusOK, vo.NewResp(e, vo.CommandExecResp{
		TransitionIds: transitionsIdList,
	}))
}
