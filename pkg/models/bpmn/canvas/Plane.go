package canvas

import (
	"fmt"

	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
)

// NewPlane ...
func NewPlane() PlaneRepository {
	return &Plane{}
}

// SetID ...
func (plane *Plane) SetID(typ string, suffix interface{}) {
	switch typ {
	case "plane":
		//plane.ID = "BPMNPlane_" + strconv.FormatInt(num, 16)
		plane.ID = fmt.Sprintf("BPMNPlane_%d", suffix)
		break
	}
}

// SetElement ...
func (plane *Plane) SetElement(typ string, suffix interface{}) {
	switch typ {
	case "process":
		plane.Element = fmt.Sprintf("Process_%s", suffix)
		break
	}
}

// SetAttrProcessElement ...
func (plane *Plane) SetAttrProcessElement(suffix string) {
	plane.Element = fmt.Sprintf("Process_%s", suffix)
}

// SetShape ...
func (plane *Plane) SetShape(num int) {
	plane.Shape = make([]Shape, num)
}

// SetEdge ...
func (plane *Plane) SetEdge(num int) {
	plane.Edge = make([]Edge, num)
}

// GetID ...
func (plane Plane) GetID() impl.STR_PTR {
	return &plane.ID
}

// GetElement ...
func (plane Plane) GetElement() impl.STR_PTR {
	return &plane.Element
}

// GetShape ...
func (plane Plane) GetShape(num int) *Shape {
	return &plane.Shape[num]
}

// GetEdge ...
func (plane Plane) GetEdge(num int) *Edge {
	return &plane.Edge[num]
}
