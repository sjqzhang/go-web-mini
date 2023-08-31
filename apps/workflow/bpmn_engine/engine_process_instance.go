package bpmn_engine

import (
	"go-web-mini/apps/workflow/spec/BPMN20"
	"go-web-mini/apps/workflow/spec/BPMN20/process_instance"
	"time"
)

type ProcessInstanceInfo struct {
	processInfo     *ProcessInfo
	instanceKey     int64
	variableContext map[string]interface{}
	createdAt       time.Time
	state           process_instance.State
	caughtEvents    []CatchEvent

	walkedElement []BPMN20.BaseElement
	currentTask   BPMN20.TServiceTask
}

func NewProcessInstanceInfo(processInfo *ProcessInfo, instanceKey int64, variableContext map[string]interface{}, createdAt time.Time, state process_instance.State, caughtEvents []CatchEvent, walkedTasks []BPMN20.BaseElement, currentTask BPMN20.TServiceTask) *ProcessInstanceInfo {
	return &ProcessInstanceInfo{processInfo: processInfo, instanceKey: instanceKey, variableContext: variableContext, createdAt: createdAt, state: state, caughtEvents: caughtEvents, walkedElement: walkedTasks, currentTask: currentTask}
}

type ProcessInstance interface {
	GetProcessInfo() *ProcessInfo
	GetInstanceKey() int64
	GetVariable(key string) string
	SetVariable(key string, value string)
	GetCreatedAt() time.Time
	GetState() process_instance.State
}

func (pii *ProcessInstanceInfo) GetProcessInfo() *ProcessInfo {
	return pii.processInfo
}

func (pii *ProcessInstanceInfo) GetInstanceKey() int64 {
	return pii.instanceKey
}

func (pii *ProcessInstanceInfo) GetVariable(key string) interface{} {
	return pii.variableContext[key]
}

func (pii *ProcessInstanceInfo) GetVariables() map[string]interface{} {
	return pii.variableContext
}

func (pii *ProcessInstanceInfo) SetVariable(key string, value string) {
	pii.variableContext[key] = value
}

func (pii *ProcessInstanceInfo) GetCreatedAt() time.Time {
	return pii.createdAt
}

// GetState returns one of [ProcessInstanceReady,ProcessInstanceActive,ProcessInstanceCompleted]
//  ┌─────┐
//  │Ready│
//  └──┬──┘
// ┌───▽──┐
// │Active│
// └───┬──┘
//┌────▽────┐
//│Completed│
//└─────────┘
func (pii *ProcessInstanceInfo) GetState() process_instance.State {
	return pii.state
}

func (pii *ProcessInstanceInfo) GetCurrentTask() BPMN20.TServiceTask {
	return pii.currentTask
}

func (pii *ProcessInstanceInfo) GetWalkedTasks() []BPMN20.BaseElement {
	return pii.walkedElement
}
