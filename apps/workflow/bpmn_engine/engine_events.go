package bpmn_engine

import (
	"fmt"
	"go-web-mini/apps/workflow/spec/BPMN20"
	"go-web-mini/apps/workflow/spec/BPMN20/activity"

	"time"
)

type MessageSubscription struct {
	ElementId          string
	ElementInstanceKey int64
	ProcessInstanceKey int64
	Name               string
	State              activity.LifecycleState
	CreatedAt          time.Time
}

type CatchEvent struct {
	Name       string
	CaughtAt   time.Time
	IsConsumed bool
}

func (state *BpmnEngineState) handleIntermediateMessageCatchEvent(id string, name string, instance *ProcessInstanceInfo) bool {
	// if auto event return true
	autos := state.GetNextAvailableAutos(instance)
	for _, auto := range autos {
		if auto.Id == id {
			return true
		}
	}

	var caughtEvent *CatchEvent
	// find first matching caught event
	for i, ce := range instance.caughtEvents {
		if ce.IsConsumed || ce.Name != name {
			continue
		}
		caughtEvent = &instance.caughtEvents[i]
	}

	commands := state.GetNextAvailableCommands(instance)
	for _, command := range commands {
		if command.GetId() == id && caughtEvent != nil {
			caughtEvent.IsConsumed = true
			return continueNextElement
		}
	}

	return false
}

func (state *BpmnEngineState) PublishEventForInstance(processInstanceKey int64, messageName string) error {
	processInstance := state.findProcessInstance(processInstanceKey)
	matchMsg := false
	for _, message := range processInstance.processInfo.definitions.Messages {
		if message.Name == messageName {
			matchMsg = true
			break
		}
	}
	if matchMsg && processInstance != nil {
		event := CatchEvent{
			CaughtAt:   time.Now(),
			Name:       messageName,
			IsConsumed: false,
		}
		processInstance.caughtEvents = append(processInstance.caughtEvents, event)
	} else {
		return fmt.Errorf("no process instance or not match event with key=%d found", processInstanceKey)
	}
	return nil
}

func (state *BpmnEngineState) findProcessInstance(processInstanceKey int64) *ProcessInstanceInfo {
	for _, pi := range state.processInstances {
		if pi.GetInstanceKey() == processInstanceKey {
			return pi
		}
	}
	return nil
}

func (state *BpmnEngineState) GetNextAvailableAutos(instance *ProcessInstanceInfo) []*BPMN20.TIntermediateCatchEvent {
	ret := make([]*BPMN20.TIntermediateCatchEvent, 0)
	outgoingAssociations := instance.currentTask.GetOutgoingAssociation()
	for _, outgoingAssociation := range outgoingAssociations {
		for _, catchEvent := range instance.processInfo.definitions.Process.IntermediateCatchEvent {
			if catchEvent.IncomingAssociation[0] != outgoingAssociation {
				continue
			}
			// Filter auto
			if catchEvent.Name != "" {
				continue
			}
			// check catchEvent's listeners is available
			isAvailable := true
			for _, extensionListener := range catchEvent.GetExtensionListeners() {
				if listener, exist := state.listeners[extensionListener.Class]; exist {
					if !listener.IsAvailable(instance, extensionListener.Fields) {
						isAvailable = false
						break
					}
				}
			}
			if isAvailable {
				ceCp := catchEvent
				ret = append(ret, &ceCp)
			}
		}
	}
	return ret
}

func (state *BpmnEngineState) GetNextAvailableCommands(instance *ProcessInstanceInfo) []*BPMN20.TIntermediateCatchEvent {
	ret := make([]*BPMN20.TIntermediateCatchEvent, 0)
	outgoingAssociations := instance.currentTask.GetOutgoingAssociation()
	for _, outgoingAssociation := range outgoingAssociations {
		for _, catchEvent := range instance.processInfo.definitions.Process.IntermediateCatchEvent {
			if catchEvent.IncomingAssociation[0] != outgoingAssociation {
				continue
			}
			// Filter auto
			if catchEvent.Name == "" {
				continue
			}
			// check catchEvent's listeners is available
			isAvailable := true
			for _, extensionListener := range catchEvent.GetExtensionListeners() {
				if listener, exist := state.listeners[extensionListener.Class]; exist {
					if !listener.IsAvailable(instance, extensionListener.Fields) {
						isAvailable = false
						break
					}
				}
			}
			if isAvailable {
				ceCp := catchEvent
				ret = append(ret, &ceCp)
			}
		}
	}
	return ret
}
