package elements

import (
	"fmt"

	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/marker"
)

// NewStartEvent ...
func NewStartEvent() StartEventRepository {
	return &StartEvent{}
}

// SetID ...
func (startEvent *StartEvent) SetID(typ string, suffix interface{}) {
	startEvent.ID = SetID(typ, suffix)
}

// SetName ...
func (startEvent *StartEvent) SetName(name string) {
	startEvent.Name = name
}

// SetOutgoing ...
func (startEvent *StartEvent) SetOutgoing(num int) {
	startEvent.Outgoing = make([]marker.Outgoing, num)
}

// GetID ...
func (startEvent StartEvent) GetID() impl.STR_PTR {
	return &startEvent.ID
}

// GetName ...
func (startEvent StartEvent) GetName() impl.STR_PTR {
	return &startEvent.Name
}

// GetOutgoing ...
func (startEvent StartEvent) GetOutgoing(num int) *marker.Outgoing {
	return &startEvent.Outgoing[num]
}

// String ...
func (startEvent StartEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", startEvent.ID, startEvent.Name)
}

// String ...
func (startEvent TStartEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", startEvent.ID, startEvent.Name)
}
