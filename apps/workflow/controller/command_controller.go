package adapter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-mini/apps/workflow/service"
	"go-web-mini/apps/workflow/vo"
	"net/http"
)

//@router /api/bpmn
type CommandController struct {

}

//@router /instance/:instance_id/commands [get]
func (ctrl *CommandController)GetAvailableCommands(c *gin.Context,req *vo.GetInstanceAvailableCommandsReq) (interface{},error) {
	ctx := c.Request.Context()
	//var req vo.GetInstanceAvailableCommandsReq
	if err := easyBind(c, &req); err != nil {
		return nil,err
	}
	instanceService := service.NewInstanceService()
	events, e := instanceService.GetAvailableCommands(ctx, req.InstanceId, req.IdentityId)
	if e != nil {
		return nil, fmt.Errorf("get available commands failed: %w", e)
	}
	commands := make([]vo.Command, len(events))
	for i, event := range events {
		commands[i] = vo.Command{
			Key:       event.Name,
			Variables: event.GetExtensionProperties(),
		}
	}
	return vo.GetInstanceAvailableCommands{Commands: commands},nil
}

//@router /instance/:instance_id/command [post]
func (ctrl *CommandController)CommandExecution(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.CommandExecution
	if err := easyBind(c, &req); err != nil {
		return
	}

	instanceService := service.NewInstanceService()
	e := instanceService.CommandExecute(ctx, req.InstanceId, req.Command, req.IdentityId, varJsonToMap(req.Variables))
	c.JSON(http.StatusOK, vo.NewResp(e, nil))
}
