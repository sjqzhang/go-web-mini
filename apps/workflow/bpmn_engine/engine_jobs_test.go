package bpmn_engine

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"

	"testing"
)

func Test_a_job_can_fail_and_keeps_the_instance_in_active_state(t *testing.T) {
	// setup
	bpmnEngine := New("name")
	process, _ := bpmnEngine.LoadFromFile("../../test-cases/simple_task.bpmn")
	bpmnEngine.AddTaskHandler("id", handler)

	instance, _ := bpmnEngine.CreateAndRunInstance(process.ProcessKey, nil)

	then.AssertThat(t, instance.state, is.EqualTo(process_instance.ACTIVE))
}

func handler(job ActivatedJob) {
	job.Fail("just because I can")
}

type blockListener struct {}

func (l blockListener) Handle(processInstanceInfo *ProcessInstanceInfo, fields []BPMN20.TExtensionExecutionListenerField) bool {
	return false
}

func (l blockListener) IsAvailable(processInstanceInfo *ProcessInstanceInfo, fields []BPMN20.TExtensionExecutionListenerField) bool {
	return false
}

type passListener struct {}

func (p passListener) Handle(processInstanceInfo *ProcessInstanceInfo, fields []BPMN20.TExtensionExecutionListenerField) bool {
	return true
}

func (p passListener) IsAvailable(processInstanceInfo *ProcessInstanceInfo, fields []BPMN20.TExtensionExecutionListenerField) bool {
	return true
}

func Test_a_job_with_listener_block(t *testing.T) {
	// setup
	bpmnEngine := New("name")
	process, _ := bpmnEngine.LoadFromFile("../../test-cases/simple_task_with_listener.bpmn")

	bpmnEngine.AddListener("test", blockListener{})
	bpmnEngine.AddTaskHandler("id", func(job ActivatedJob) {
		job.Complete()
	})

	instance, _ := bpmnEngine.CreateAndRunInstance(process.ProcessKey, nil)

	then.AssertThat(t, instance.state, is.EqualTo(process_instance.ACTIVE))
}

func Test_a_job_with_listener_pass(t *testing.T) {
	// setup
	bpmnEngine := New("name")
	process, _ := bpmnEngine.LoadFromFile("../../test-cases/simple_task_with_listener.bpmn")

	bpmnEngine.AddListener("test", passListener{})
	bpmnEngine.AddTaskHandler("id", func(job ActivatedJob) {
		job.Complete()
	})

	instance, _ := bpmnEngine.CreateAndRunInstance(process.ProcessKey, map[string]interface{}{"username": "abc"})

	then.AssertThat(t, instance.state, is.EqualTo(process_instance.COMPLETED))
}

