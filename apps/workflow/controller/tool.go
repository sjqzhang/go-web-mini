package adapter

import (

	"github.com/gin-gonic/gin"
	"go-web-mini/apps/workflow/spec/BPMN20"
	"go-web-mini/apps/workflow/vo"
	"net/http"
)

func easyBind(c *gin.Context, req interface{}) (err error) {
	if err = c.ShouldBindUri(req); err != nil {
		c.JSON(http.StatusBadRequest, &vo.CommResp{RetCode: http.StatusBadRequest, Message: err.Error()})
		return err
	}

	if err = c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &vo.CommResp{RetCode: http.StatusBadRequest, Message: err.Error()})
		return err
	}
	return nil
}

func varJsonToMap(vars []vo.Variables) map[string]interface{} {
	res := make(map[string]interface{}, 0)
	for _, v := range vars {
		res[v.Name] = v.Value
	}
	return res
}

func convertServiceTasks(tasks []BPMN20.TServiceTask) []vo.ServiceTasks {
	var serviceTasks []vo.ServiceTasks
	for _, task := range tasks {
		serviceTasks = append(serviceTasks, vo.ServiceTasks{TaskId: task.Id, TaskName: task.Name})
	}
	return serviceTasks
}

func convertVariables(varMap map[string]interface{}) []vo.Variables {
	var variables []vo.Variables
	for k, v := range varMap {
		variables = append(variables, vo.Variables{Name: k, Value: v.(string)})
	}
	return variables
}
