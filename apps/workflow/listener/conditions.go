package listener

import (

	"github.com/antonmedv/expr"
	"github.com/sirupsen/logrus"
	"go-web-mini/apps/workflow/bpmn_engine"
	"go-web-mini/apps/workflow/spec/BPMN20"
	"strings"
)

func evaluateExpression(expression string, variableContext map[string]interface{}) (interface{}, error) {
	expression = strings.TrimSpace(expression)
	expression = strings.TrimPrefix(expression, "=")
	return expr.Eval(expression, variableContext)
}

type ConditionsExpressionChecker struct {}

func (c ConditionsExpressionChecker) Handle(processInstanceInfo *bpmn_engine.ProcessInstanceInfo, fields []BPMN20.TExtensionExecutionListenerField) bool {
	expression := ""
	for _, field := range fields {
		if field.Name == "expression" {
			expression = field.Value
		}
	}
	if expression == "" {
		return false
	}
	variables := processInstanceInfo.GetVariables()
	result, err := evaluateExpression(expression, variables)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"expression": expression,
			"variables": variables,
		}).Error("evaluateExpression err")
		return false
	}
	return result.(bool)
}

func (c ConditionsExpressionChecker) IsAvailable(processInstanceInfo *bpmn_engine.ProcessInstanceInfo, fields []BPMN20.TExtensionExecutionListenerField) bool {
	return c.Handle(processInstanceInfo, fields)
}


