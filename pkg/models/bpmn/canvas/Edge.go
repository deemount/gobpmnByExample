package canvas

import (
	"fmt"

	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
)

// NewEdge ...
func NewEdge() EdgeRepository {
	return &Edge{}
}

// SetID ...
func (edge *Edge) SetID(typ string, suffix interface{}) {
	switch typ {
	case "flow":
		edge.ID = fmt.Sprintf("Flow_%s_di", suffix)
	}
}

// SetElement ...
func (edge *Edge) SetElement(typ string, suffix interface{}) {
	switch typ {
	case "flow":
		edge.Element = fmt.Sprintf("Flow_%s", suffix)
	}
}

// SetWaypoint ...
func (edge *Edge) SetWaypoint() {
	edge.Waypoint = make([]Waypoint, 2)
}

// SetWaypoints ...
func (edge *Edge) SetWaypoints(num int) {
	edge.Waypoint = make([]Waypoint, num)
}

// GetID ...
func (edge Edge) GetID() impl.STR_PTR {
	return &edge.ID
}

// GetElement ...
func (edge Edge) GetElement() impl.STR_PTR {
	return &edge.Element
}

// GetWaypoint ...
func (edge Edge) GetWaypoint(num int) *Waypoint {
	return &edge.Waypoint[num]
}
