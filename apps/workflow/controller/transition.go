package controller

import (

	"github.com/gin-gonic/gin"
	"go-web-mini/apps/workflow/service"
	"go-web-mini/apps/workflow/vo"
	"net/http"
)

//@Description transition信息
//@router /bpmn/api
type TransitionController struct{

}

//@description 获取实例的transition信息
//@router /transitions/ids [post]
func (cc *TransitionController)GetInstanceTransitionsByIds(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.GetInstanceTransitionsByIds
	if err := easyBind(c, &req); err != nil {
		return
	}
	instService := service.NewInstanceService()
	transitions, e := instService.GetTransitionsByIds(ctx, req.Ids)
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
