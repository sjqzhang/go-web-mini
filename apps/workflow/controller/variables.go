package adapter

import (

	"github.com/gin-gonic/gin"
	"go-web-mini/apps/workflow/service"
	"go-web-mini/apps/workflow/vo"
	"net/http"
)

func GetInstanceVariables(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.GetInstanceVariablesReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	variablesService := service.NewVariablesService()
	name, value, e := variablesService.GetInstanceVariables(ctx, req.InstanceId, req.Name)
	if e != nil {
		c.JSON(http.StatusOK, vo.NewResp(e, nil))
		return
	}
	c.JSON(http.StatusOK, vo.NewResp(nil, vo.GetVariablesResp{Name: name, Value: value}))
}

func SetInstanceVariables(c *gin.Context) {
	ctx := c.Request.Context()
	var req vo.SetInstanceVariablesReq
	if err := easyBind(c, &req); err != nil {
		return
	}
	variablesService := service.NewVariablesService()
	e := variablesService.SetInstanceVariables(ctx, req.InstanceId, req.Name, req.Value)
	c.JSON(http.StatusOK, vo.NewResp(e, nil))
}
