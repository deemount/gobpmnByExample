package elements

import (
	"fmt"

	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/marker"
)

// NewEndEvent ...
func NewEndEvent() EndEventRepository {
	return &EndEvent{}
}

// SetID ...
func (endEvent *EndEvent) SetID(typ string, suffix interface{}) {
	endEvent.ID = SetID(typ, suffix)
}

// SetName ...
func (endEvent *EndEvent) SetName(name string) {
	endEvent.Name = name
}

// SetIncoming ...
func (endEvent *EndEvent) SetIncoming(num int) {
	endEvent.Incoming = make([]marker.Incoming, num)
}

// GetID ...
func (endEvent EndEvent) GetID() impl.STR_PTR {
	return &endEvent.ID
}

// GetName ...
func (endEvent EndEvent) GetName() impl.STR_PTR {
	return &endEvent.Name
}

// GetIncoming ...
func (endEvent EndEvent) GetIncoming(num int) *marker.Incoming {
	return &endEvent.Incoming[num]
}

// String ...
func (endEvent EndEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", endEvent.ID, endEvent.Name)
}

// String ...
func (endEvent TEndEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", endEvent.ID, endEvent.Name)
}
