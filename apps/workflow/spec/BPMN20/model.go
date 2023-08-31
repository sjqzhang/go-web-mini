package BPMN20

import (
	"encoding/xml"
)

type TDefinitions struct {
	XMLName            xml.Name   `xml:"definitions"`
	Id                 string     `xml:"id,attr"`
	Name               string     `xml:"name,attr"`
	TargetNamespace    string     `xml:"targetNamespace,attr"`
	ExpressionLanguage string     `xml:"expressionLanguage,attr"`
	TypeLanguage       string     `xml:"typeLanguage,attr"`
	Exporter           string     `xml:"exporter,attr"`
	ExporterVersion    string     `xml:"exporterVersion,attr"`
	Process            TProcess   `xml:"process"`
	Messages           []TMessage `xml:"message"`
}

type TProcess struct {
	XMLName                      xml.Name                  `xml:"process"`
	Id                           string                    `xml:"id,attr"`
	Name                         string                    `xml:"name,attr"`
	ProcessType                  string                    `xml:"processType,attr"`
	IsClosed                     bool                      `xml:"isClosed,attr"`
	IsExecutable                 bool                      `xml:"isExecutable,attr"`
	DefinitionalCollaborationRef string                    `xml:"definitionalCollaborationRef,attr"`
	StartEvents                  []TStartEvent             `xml:"startEvent"`
	EndEvents                    []TEndEvent               `xml:"endEvent"`
	SequenceFlows                []TSequenceFlow           `xml:"sequenceFlow"`
	ServiceTasks                 []TServiceTask            `xml:"serviceTask"`
	ParallelGateway              []TParallelGateway        `xml:"parallelGateway"`
	ExclusiveGateway             []TExclusiveGateway       `xml:"exclusiveGateway"`
	IntermediateCatchEvent       []TIntermediateCatchEvent `xml:"intermediateCatchEvent"`
	EventBasedGateway            []TEventBasedGateway      `xml:"eventBasedGateway"`
}

type TSequenceFlow struct {
	XMLName             xml.Name      `xml:"sequenceFlow"`
	Id                  string        `xml:"id,attr"`
	Name                string        `xml:"name,attr"`
	SourceRef           string        `xml:"sourceRef,attr"`
	TargetRef           string        `xml:"targetRef,attr"`
	ConditionExpression []TExpression `xml:"conditionExpression"`
}

type TExpression struct {
	XMLName xml.Name `xml:"conditionExpression"`
	Text    string   `xml:",innerxml"`
}

type TStartEvent struct {
	XMLName             xml.Name           `xml:"startEvent"`
	Id                  string             `xml:"id,attr"`
	Name                string             `xml:"name,attr"`
	IsInterrupting      bool               `xml:"isInterrupting,attr"`
	ParallelMultiple    bool               `xml:"parallelMultiple,attr"`
	IncomingAssociation []string           `xml:"incoming"`
	OutgoingAssociation []string           `xml:"outgoing"`
	ExtensionElements   TExtensionElements `xml:"extensionElements"`
}

type TEndEvent struct {
	XMLName             xml.Name           `xml:"endEvent"`
	Id                  string             `xml:"id,attr"`
	Name                string             `xml:"name,attr"`
	IncomingAssociation []string           `xml:"incoming"`
	OutgoingAssociation []string           `xml:"outgoing"`
	ExtensionElements   TExtensionElements `xml:"extensionElements"`
}

type TServiceTask struct {
	XMLName             xml.Name           `xml:"serviceTask"`
	Id                  string             `xml:"id,attr"`
	Name                string             `xml:"name,attr"`
	Default             string             `xml:"default,attr"`
	CompletionQuantity  int                `xml:"completionQuantity,attr"`
	IsForCompensation   bool               `xml:"isForCompensation,attr"`
	OperationRef        string             `xml:"operationRef,attr"`
	Implementation      string             `xml:"implementation,attr"`
	IncomingAssociation []string           `xml:"incoming"`
	OutgoingAssociation []string           `xml:"outgoing"`
	ExtensionElements   TExtensionElements `xml:"extensionElements"`
}

type TParallelGateway struct {
	XMLName             xml.Name           `xml:"parallelGateway"`
	Id                  string             `xml:"id,attr"`
	Name                string             `xml:"name,attr"`
	IncomingAssociation []string           `xml:"incoming"`
	OutgoingAssociation []string           `xml:"outgoing"`
	ExtensionElements   TExtensionElements `xml:"extensionElements"`
}

type TExclusiveGateway struct {
	XMLName             xml.Name           `xml:"exclusiveGateway"`
	Id                  string             `xml:"id,attr"`
	Name                string             `xml:"name,attr"`
	IncomingAssociation []string           `xml:"incoming"`
	OutgoingAssociation []string           `xml:"outgoing"`
	ExtensionElements   TExtensionElements `xml:"extensionElements"`
}

type TIntermediateCatchEvent struct {
	XMLName                xml.Name                `xml:"intermediateCatchEvent"`
	Id                     string                  `xml:"id,attr"`
	Name                   string                  `xml:"name,attr"`
	IncomingAssociation    []string                `xml:"incoming"`
	OutgoingAssociation    []string                `xml:"outgoing"`
	MessageEventDefinition TMessageEventDefinition `xml:"messageEventDefinition"`
	TimerEventDefinition   TTimerEventDefinition   `xml:"timerEventDefinition"`
	ParallelMultiple       bool                    `xml:"parallelMultiple"`
	ExtensionElements      TExtensionElements      `xml:"extensionElements"`
}

type TEventBasedGateway struct {
	XMLName             xml.Name           `xml:"eventBasedGateway"`
	Id                  string             `xml:"id,attr"`
	Name                string             `xml:"name,attr"`
	IncomingAssociation []string           `xml:"incoming"`
	OutgoingAssociation []string           `xml:"outgoing"`
	ExtensionElements   TExtensionElements `xml:"extensionElements"`
}

type TMessageEventDefinition struct {
	XMLName    xml.Name `xml:"messageEventDefinition"`
	Id         string   `xml:"id,attr"`
	MessageRef string   `xml:"messageRef,attr"`
}

type TTimerEventDefinition struct {
	XMLName      xml.Name      `xml:"timerEventDefinition"`
	Id           string        `xml:"id,attr"`
	TimeDuration TTimeDuration `xml:"timeDuration"`
}

type TMessage struct {
	XMLName xml.Name `xml:"message"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type TTimeDuration struct {
	XMLName xml.Name `xml:"timeDuration"`
	XMLText string   `xml:",innerxml"`
}

type TExtensionElements struct {
	XMLName            xml.Name                      `xml:"extensionElements"`
	Properties         TExtensionProperties          `xml:"properties"`
	ExecutionListeners []TExtensionExecutionListener `xml:"executionListener"`
}

type TExtensionProperties struct {
	XMLName    xml.Name             `xml:"properties"`
	Properties []TExtensionProperty `xml:"property"`
}

type TExtensionProperty struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type TExtensionExecutionListener struct {
	XMLName xml.Name `xml:"executionListener"`
	Class   string   `xml:"class,attr"`
	Event   string   `xml:"event,attr"`
	Fields  []TExtensionExecutionListenerField `xml:"field"`
}

type TExtensionExecutionListenerField struct {
	XMLName xml.Name `xml:"field"`
	Name string `xml:"name,attr"`
	Value string `xml:"string"`
}
