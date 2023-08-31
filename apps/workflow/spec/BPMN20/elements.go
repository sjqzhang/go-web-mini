package BPMN20

type ElementType string

const (
	StartEvent             ElementType = "START_EVENT"
	EndEvent               ElementType = "END_EVENT"
	ServiceTask            ElementType = "SERVICE_TASK"
	ParallelGateway        ElementType = "PARALLEL_GATEWAY"
	ExclusiveGateway       ElementType = "EXCLUSIVE_GATEWAY"
	IntermediateCatchEvent ElementType = "INTERMEDIATE_CATCH_EVENT"
	EventBasedGateway      ElementType = "EVENT_BASED_GATEWAY"
)

type BaseElement interface {
	GetId() string
	GetName() string
	GetIncomingAssociation() []string
	GetOutgoingAssociation() []string
	GetType() ElementType
	GetExtensionProperties() []string
	GetExtensionListeners() []TExtensionExecutionListener
}

func (startEvent TStartEvent) GetId() string {
	return startEvent.Id
}

func (startEvent TStartEvent) GetName() string {
	return startEvent.Name
}

func (startEvent TStartEvent) GetIncomingAssociation() []string {
	return startEvent.IncomingAssociation
}

func (startEvent TStartEvent) GetOutgoingAssociation() []string {
	return startEvent.OutgoingAssociation
}

func (startEvent TStartEvent) GetType() ElementType {
	return StartEvent
}

func (startEvent TStartEvent) GetExtensionProperties() []string {
	names := make([]string, len(startEvent.ExtensionElements.Properties.Properties))
	for i, property := range startEvent.ExtensionElements.Properties.Properties {
		names[i] = property.Name
	}
	return names
}

func (startEvent TStartEvent) GetExtensionListeners() []TExtensionExecutionListener {
	return startEvent.ExtensionElements.ExecutionListeners
}

func (endEvent TEndEvent) GetId() string {
	return endEvent.Id
}

func (endEvent TEndEvent) GetName() string {
	return endEvent.Name
}

func (endEvent TEndEvent) GetIncomingAssociation() []string {
	return endEvent.IncomingAssociation
}

func (endEvent TEndEvent) GetOutgoingAssociation() []string {
	return endEvent.OutgoingAssociation
}

func (endEvent TEndEvent) GetType() ElementType {
	return EndEvent
}

func (endEvent TEndEvent) GetExtensionProperties() []string {
	names := make([]string, len(endEvent.ExtensionElements.Properties.Properties))
	for i, property := range endEvent.ExtensionElements.Properties.Properties {
		names[i] = property.Name
	}
	return names
}

func (endEvent TEndEvent) GetExtensionListeners() []TExtensionExecutionListener {
	return endEvent.ExtensionElements.ExecutionListeners
}

func (serviceTask TServiceTask) GetId() string {
	return serviceTask.Id
}

func (serviceTask TServiceTask) GetName() string {
	return serviceTask.Name
}

func (serviceTask TServiceTask) GetIncomingAssociation() []string {
	return serviceTask.IncomingAssociation
}

func (serviceTask TServiceTask) GetOutgoingAssociation() []string {
	return serviceTask.OutgoingAssociation
}

func (serviceTask TServiceTask) GetType() ElementType {
	return ServiceTask
}

func (serviceTask TServiceTask) GetExtensionProperties() []string {
	names := make([]string, len(serviceTask.ExtensionElements.Properties.Properties))
	for i, property := range serviceTask.ExtensionElements.Properties.Properties {
		names[i] = property.Name
	}
	return names
}

func (serviceTask TServiceTask) GetExtensionListeners() []TExtensionExecutionListener {
	return serviceTask.ExtensionElements.ExecutionListeners
}

func (parallelGateway TParallelGateway) GetId() string {
	return parallelGateway.Id
}

func (parallelGateway TParallelGateway) GetName() string {
	return parallelGateway.Name
}

func (parallelGateway TParallelGateway) GetIncomingAssociation() []string {
	return parallelGateway.IncomingAssociation
}

func (parallelGateway TParallelGateway) GetOutgoingAssociation() []string {
	return parallelGateway.OutgoingAssociation
}

func (parallelGateway TParallelGateway) GetType() ElementType {
	return ParallelGateway
}

func (parallelGateway TParallelGateway) GetExtensionProperties() []string {
	names := make([]string, len(parallelGateway.ExtensionElements.Properties.Properties))
	for i, property := range parallelGateway.ExtensionElements.Properties.Properties {
		names[i] = property.Name
	}
	return names
}

func (parallelGateway TParallelGateway) GetExtensionListeners() []TExtensionExecutionListener {
	return parallelGateway.ExtensionElements.ExecutionListeners
}

func (exclusiveGateway TExclusiveGateway) GetId() string {
	return exclusiveGateway.Id
}

func (exclusiveGateway TExclusiveGateway) GetName() string {
	return exclusiveGateway.Name
}

func (exclusiveGateway TExclusiveGateway) GetIncomingAssociation() []string {
	return exclusiveGateway.IncomingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetOutgoingAssociation() []string {
	return exclusiveGateway.OutgoingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetType() ElementType {
	return ExclusiveGateway
}

func (exclusiveGateway TExclusiveGateway) GetExtensionProperties() []string {
	names := make([]string, len(exclusiveGateway.ExtensionElements.Properties.Properties))
	for i, property := range exclusiveGateway.ExtensionElements.Properties.Properties {
		names[i] = property.Name
	}
	return names
}

func (exclusiveGateway TExclusiveGateway) GetExtensionListeners() []TExtensionExecutionListener {
	return exclusiveGateway.ExtensionElements.ExecutionListeners
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetId() string {
	return intermediateCatchEvent.Id
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetName() string {
	return intermediateCatchEvent.Name
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetIncomingAssociation() []string {
	return intermediateCatchEvent.IncomingAssociation
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetOutgoingAssociation() []string {
	return intermediateCatchEvent.OutgoingAssociation
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetType() ElementType {
	return IntermediateCatchEvent
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetExtensionProperties() []string {
	names := make([]string, len(intermediateCatchEvent.ExtensionElements.Properties.Properties))
	for i, property := range intermediateCatchEvent.ExtensionElements.Properties.Properties {
		names[i] = property.Name
	}
	return names
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetExtensionListeners() []TExtensionExecutionListener {
	return intermediateCatchEvent.ExtensionElements.ExecutionListeners
}

func (eventBasedGateway TEventBasedGateway) GetId() string {
	return eventBasedGateway.Id
}

func (eventBasedGateway TEventBasedGateway) GetName() string {
	return eventBasedGateway.Name
}

func (eventBasedGateway TEventBasedGateway) GetIncomingAssociation() []string {
	return eventBasedGateway.IncomingAssociation
}

func (eventBasedGateway TEventBasedGateway) GetOutgoingAssociation() []string {
	return eventBasedGateway.OutgoingAssociation
}

func (eventBasedGateway TEventBasedGateway) GetType() ElementType {
	return EventBasedGateway
}

func (eventBasedGateway TEventBasedGateway) GetExtensionProperties() []string {
	names := make([]string, len(eventBasedGateway.ExtensionElements.Properties.Properties))
	for i, property := range eventBasedGateway.ExtensionElements.Properties.Properties {
		names[i] = property.Name
	}
	return names
}

func (eventBasedGateway TEventBasedGateway) GetExtensionListeners() []TExtensionExecutionListener {
	return eventBasedGateway.ExtensionElements.ExecutionListeners
}
