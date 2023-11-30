package canvas

import (
	"fmt"

	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
)

// NewShape ...
func NewShape() ShapeRepository {
	return &Shape{}
}

// SetID ...
func (shape *Shape) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		shape.ID = fmt.Sprintf("Activity_%s_di", suffix)
		break
	case "event":
		shape.ID = fmt.Sprintf("Event_%s_di", suffix)
		break
	case "startevent":
		shape.ID = fmt.Sprintf("_BPMNShape_StartEvent_%v", suffix)
		break
	}
}

// SetElement ...
func (shape *Shape) SetElement(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		shape.Element = fmt.Sprintf("Activity_%s", suffix)
		break
	case "event":
		shape.Element = fmt.Sprintf("Event_%s", suffix)
		break
	case "startevent":
		shape.Element = fmt.Sprintf("StartEvent_%v", suffix)
		break
	}
}

// SetBounds ...
func (shape *Shape) SetBounds() {
	shape.Bounds = make([]Bounds, 1)
}

// GetID ...
func (shape Shape) GetID() impl.STR_PTR {
	return &shape.ID
}

// GetElement ...
func (shape Shape) GetElement() impl.STR_PTR {
	return &shape.Element
}

// GetBounds ...
func (shape Shape) GetBounds() *Bounds {
	return &shape.Bounds[0]
}
