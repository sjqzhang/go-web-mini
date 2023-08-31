package listener

import (
	"go-web-mini/apps/workflow/bpmn_engine"
	"go-web-mini/apps/workflow/spec/BPMN20"
	"strings"
)

type OrApprovalListener struct {}

func (o OrApprovalListener) Handle(processInstanceInfo *bpmn_engine.ProcessInstanceInfo, fields []BPMN20.TExtensionExecutionListenerField) bool {
	approversParam, approverParam := "", ""
	for _, field := range fields {
		if field.Name == "approver" {
			approverParam = field.Value
		}
		if field.Name == "approvers" {
			approversParam = field.Value
		}
	}
	if approversParam == "" || approverParam == "" {
		return false
	}
	if processInstanceInfo.GetVariable(approverParam) == nil || processInstanceInfo.GetVariable(approversParam) == nil {
		return false
	}
	approver := processInstanceInfo.GetVariable(approverParam).(string)
	for _, allowedApprover := range strings.Split(processInstanceInfo.GetVariable(approversParam).(string), ",") {
		if approver == allowedApprover {
			return true
		}
	}
	return false
}

func (o OrApprovalListener) IsAvailable(processInstanceInfo *bpmn_engine.ProcessInstanceInfo, fields []BPMN20.TExtensionExecutionListenerField) bool {
	return o.Handle(processInstanceInfo, fields)
}

