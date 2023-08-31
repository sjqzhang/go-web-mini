package bpmn_engine

import (
	"github.com/bwmarrin/snowflake"
	"go-web-mini/apps/workflow/spec/BPMN20"
)

type ProcessInfo struct {
	BpmnProcessId string // The ID as defined in the BPMN file
	Version       int32  // A version of the process, default=1, incremented, when another process with the same ID is loaded
	ProcessKey    int64  // The engines key for this given process with version

	definitions   BPMN20.TDefinitions // parsed file content
	checksumBytes [16]byte            // internal checksum to identify different versions
}

func (p *ProcessInfo) Definitions() BPMN20.TDefinitions {
	return p.definitions
}

type BpmnEngineListener interface {
	Handle(processInstanceInfo *ProcessInstanceInfo, fields []BPMN20.TExtensionExecutionListenerField) bool
	IsAvailable(processInstanceInfo *ProcessInstanceInfo, fields []BPMN20.TExtensionExecutionListenerField) bool
}

type BpmnEngineState struct {
	name                 string
	processes            []ProcessInfo
	processInstances     []*ProcessInstanceInfo
	jobs                 []*job
	timers               []*Timer
	// TIP: scheduledFlows just for handleParallelGateway, not compatible now.
	scheduledFlows       []string
	handlers             map[string]func(job ActivatedJob)
	listeners            map[string]BpmnEngineListener
	snowflake            *snowflake.Node
}

// GetProcessInstances returns a list of instance information.
func (state *BpmnEngineState) GetProcessInstances() []*ProcessInstanceInfo {
	return state.processInstances
}

func (state *BpmnEngineState) SetProcessInstance(instance *ProcessInstanceInfo) {
	state.processInstances = append(state.processInstances, instance)
}

// FindProcessInstanceById searches for a give processInstanceKey
// and returns the corresponding ProcessInstanceInfo otherwise nil
func (state *BpmnEngineState) FindProcessInstanceById(processInstanceKey int64) *ProcessInstanceInfo {
	for _, instance := range state.processInstances {
		if instance.instanceKey == processInstanceKey {
			return instance
		}
	}
	return nil
}

// GetName returns the name of the engine, only useful in case you control multiple ones
func (state *BpmnEngineState) GetName() string {
	return state.name
}
