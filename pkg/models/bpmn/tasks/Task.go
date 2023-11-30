package tasks

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

// NewTask ...
func NewTask() TaskRepository {
	return &Task{}
}

// SetID ...
func (task *Task) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		task.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		task.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (task *Task) SetName(name string) {
	task.Name = name
}

// SetIncoming ...
func (task *Task) SetIncoming(num int) {
	task.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (task *Task) SetOutgoing(num int) {
	task.Outgoing = make([]marker.Outgoing, num)
}

// GetID ...
func (task Task) GetID() impl.STR_PTR {
	return &task.ID
}

// GetName ...
func (task Task) GetName() impl.STR_PTR {
	return &task.Name
}

// GetIncoming ...
func (task Task) GetIncoming(num int) *marker.Incoming {
	return &task.Incoming[num]
}

// GetOutgoing ...
func (task Task) GetOutgoing(num int) *marker.Outgoing {
	return &task.Outgoing[num]
}

// String ...
func (task Task) String() string {
	return fmt.Sprintf("id=%v, name=%v", task.ID, task.Name)
}
