package adapter

import (


	"github.com/gin-gonic/gin"
	"go-web-mini/apps/workflow/service"
	"go-web-mini/apps/workflow/vo"
	"net/http"
)

func CreateInstance(c *gin.Context) {
	ctx := c.Request.Context()
	var createReq vo.CreateInstance
	if err := easyBind(c, &createReq); err != nil {
		return
	}
	instService := service.NewInstanceService()
	instanceId, e := instService.CreateInstance(ctx, createReq.SchemeCode, varJsonToMap(createReq.Variables))
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, vo.CreateInstanceResp{InstanceID: instanceId}))
}

func GetInstance(c *gin.Context) {
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
	variables := convertVariables(instanceInfo.GetVariables())
	resp := vo.GetInstanceResp{
		InstanceID:    instanceInfo.Id,
		SchemeCode:    instanceInfo.SchemeCode,
		State:         instanceInfo.State,
		CurrentTaskID: instanceInfo.CurrentTaskId,
		ServiceTasks:  serviceTasks,
		Variables:     variables,
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, resp))
}

func GetInstanceTransitions(c *gin.Context) {
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
