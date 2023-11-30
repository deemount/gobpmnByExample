package process

import (
	"fmt"

	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/events"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/events/elements"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/flow"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/tasks"
)

// NewProcess ...
func NewProcess() ProcessRepository {
	return &Process{}
}

// SetID ...
func (process *Process) SetID(typ string, suffix interface{}) {
	switch typ {
	case "process":
		process.ID = fmt.Sprintf("Process_%v", suffix)
		break
	}
}

// SetName ...
func (process *Process) SetName(name string) {
	process.Name = name
}

// SetIsExecutable ...
func (process *Process) SetIsExecutable(isExec bool) {
	process.IsExecutable = isExec
}

// SetStartEvent ...
func (process *Process) SetStartEvent(num int) {
	process.StartEvent = make([]elements.StartEvent, num)
}

// SetEndEvent ...
func (process *Process) SetEndEvent(num int) {
	process.EndEvent = make(events.END_EVENT_SLC, num)
}

// SetTask ...
func (process *Process) SetTask(num int) {
	process.Task = make([]tasks.Task, num)
}

// SetSequenceFlow ...
func (process *Process) SetSequenceFlow(num int) {
	process.SequenceFlow = make([]flow.SequenceFlow, num)
}

// GetID ...
func (process Process) GetID() impl.STR_PTR {
	return &process.ID
}

// GetName ...
func (process Process) GetName() impl.STR_PTR {
	return &process.Name
}

// GetIsExecutable ...
func (process Process) GetIsExecutable() *bool {
	return &process.IsExecutable
}

// GetStartEvent ...
func (process Process) GetStartEvent(num int) *elements.StartEvent {
	return &process.StartEvent[num]
}

// GetEndEvent ...
func (process Process) GetEndEvent(num int) events.END_EVENT_PTR {
	return &process.EndEvent[num]
}

// GetTask ...
func (process Process) GetTask(num int) tasks.TASK_PTR {
	return &process.Task[num]
}

// GetSequenceFlow ...
func (process Process) GetSequenceFlow(num int) *flow.SequenceFlow {
	return &process.SequenceFlow[num]
}
