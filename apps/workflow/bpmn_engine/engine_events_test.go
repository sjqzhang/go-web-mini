package bpmn_engine

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

func Test_creating_a_process_sets_state_to_READY(t *testing.T) {
	// setup
	bpmnEngine := New("name")

	// given
	process, _ := bpmnEngine.LoadFromFile("../../test-cases/message-intermediate-catch-event.bpmn")

	// when
	pi, _ := bpmnEngine.CreateInstance(process.ProcessKey, nil)
	// then
	then.AssertThat(t, pi.GetState(), is.EqualTo(process_instance.READY))
}

func Test_running_a_process_sets_state_to_ACTIVE(t *testing.T) {
	// setup
	bpmnEngine := New("name")

	// given
	process, _ := bpmnEngine.LoadFromFile("../../test-cases/message-intermediate-catch-event.bpmn")

	// when
	pi, _ := bpmnEngine.CreateInstance(process.ProcessKey, nil)
	procInst, _ := bpmnEngine.RunOrContinueInstance(pi.GetInstanceKey())

	// then
	then.AssertThat(t, pi.GetState(), is.EqualTo(process_instance.ACTIVE).
		Reason("Since the BPMN contains an intermediate catch event, the process instance must be active and can't complete."))
	then.AssertThat(t, procInst.GetState(), is.EqualTo(process_instance.ACTIVE))
}

func Test_IntermediateCatchEvent_received_message_completes_the_instance(t *testing.T) {
	// setup
	bpmnEngine := New("name")

	// given
	process, _ := bpmnEngine.LoadFromFile("../../test-cases/message-intermediate-catch-event.bpmn")
	pi, _ := bpmnEngine.CreateAndRunInstance(process.ProcessKey, nil)

	// when
	bpmnEngine.PublishEventForInstance(pi.GetInstanceKey(), "globalMsgRef")
	bpmnEngine.RunOrContinueInstance(pi.GetInstanceKey())

	// then
	then.AssertThat(t, pi.GetState(), is.EqualTo(process_instance.ACTIVE))
}

func Test_IntermediateCatchEvent_message_can_be_published_before_running_the_instance(t *testing.T) {
	// setup
	bpmnEngine := New("name")

	// given
	process, _ := bpmnEngine.LoadFromFile("../../test-cases/message-intermediate-catch-event.bpmn")
	pi, _ := bpmnEngine.CreateInstance(process.ProcessKey, nil)

	// when
	bpmnEngine.PublishEventForInstance(pi.GetInstanceKey(), "event-1")
	bpmnEngine.RunOrContinueInstance(pi.GetInstanceKey())

	// then
	then.AssertThat(t, pi.GetState(), is.EqualTo(process_instance.ACTIVE))
}

// func Test_Having_IntermediateCatchEvent_and_ServiceTask_in_parallel_the_process_state_is_maintained(t *testing.T) {
// 	// setup
// 	bpmnEngine := New("name")
// 	cp := CallPath{}
//
// 	// given
// 	process, _ := bpmnEngine.LoadFromFile("../../test-cases/message-intermediate-catch-event-and-parallel-tasks.bpmn")
// 	instance, _ := bpmnEngine.CreateInstance(process.ProcessKey, nil)
// 	bpmnEngine.AddTaskHandler("task-1", cp.CallPathHandler)
// 	bpmnEngine.AddTaskHandler("task-2", cp.CallPathHandler)
//
// 	// when
// 	bpmnEngine.RunOrContinueInstance(instance.GetInstanceKey())
//
// 	// then
// 	then.AssertThat(t, instance.GetState(), is.EqualTo(process_instance.ACTIVE))
//
// 	// when
// 	bpmnEngine.PublishEventForInstance(instance.GetInstanceKey(), "event-1")
// 	bpmnEngine.RunOrContinueInstance(instance.GetInstanceKey())
//
// 	// then
// 	then.AssertThat(t, cp.CallPath, is.EqualTo("task-2,task-1"))
// 	then.AssertThat(t, instance.GetState(), is.EqualTo(process_instance.COMPLETED))
// }

// func Test_two(t *testing.T) {
// 	// setup
// 	bpmnEngine := New("name")
// 	cp := CallPath{}
//
// 	// given
// 	process, _ := bpmnEngine.LoadFromFile("../../test-cases/two-tasks-shared-message-event2.bpmn")
// 	instance, _ := bpmnEngine.CreateInstance(process.ProcessKey, nil)
// 	bpmnEngine.AddTaskHandler("task-a", cp.CallPathHandler)
// 	bpmnEngine.AddTaskHandler("task-b", cp.CallPathHandler)
// 	bpmnEngine.AddListener("orApprovalHandler", func(processInstanceInfo *ProcessInstanceInfo, fields []BPMN20.TExtensionExecutionListenerField) bool {
// 		if processInstanceInfo.GetVariable("approver") == nil || processInstanceInfo.GetVariable("approvers") == nil {
// 			return false
// 		}
// 		approver := processInstanceInfo.GetVariable("approver").(string)
// 		for _, allowedApprover := range strings.Split(processInstanceInfo.GetVariable("approvers").(string), ",") {
// 			if approver == allowedApprover {
// 				return true
// 			}
// 		}
// 		return false
// 	})
//
// 	// when
// 	bpmnEngine.PublishEventForInstance(instance.GetInstanceKey(), "A")
// 	instance.SetVariable("approvers", "junjie.yu@shopee.com,jingyu.fu@shopee.com")
// 	instance.SetVariable("approver", "junjie.yu@shopee.com")
//
// 	bpmnEngine.RunOrContinueInstance(instance.GetInstanceKey())
//
// 	// then
// 	then.AssertThat(t, cp.CallPath, is.EqualTo("task-a"))
// 	then.AssertThat(t, instance.GetState(), is.EqualTo(process_instance.ACTIVE))
//
// 	bpmnEngine.PublishEventForInstance(instance.GetInstanceKey(), "B")
// 	bpmnEngine.RunOrContinueInstance(instance.GetInstanceKey())
// 	then.AssertThat(t, cp.CallPath, is.EqualTo("task-a,task-b"))
// 	then.AssertThat(t, instance.GetState(), is.EqualTo(process_instance.COMPLETED))
// }
