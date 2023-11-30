package canvas

import (
	"fmt"

	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
)

// NewDiagram ...
func NewDiagram() DiagramRepository {
	return &Diagram{}
}

// SetID ...
func (diagram *Diagram) SetID(typ string, suffix interface{}) {
	switch typ {
	case "diagram":
		//diagram.ID = "BPMNDiagram_" + strconv.FormatInt(num, 16)
		diagram.ID = fmt.Sprintf("BPMNDiagram_%v", suffix)
		break
	}
}

// SetPlane ...
func (diagram *Diagram) SetPlane() {
	diagram.Plane = make([]Plane, 1)
}

// GetID ...
func (diagram Diagram) GetID() impl.STR_PTR {
	return &diagram.ID
}

// GetPlane ...
func (diagram Diagram) GetPlane() *Plane {
	return &diagram.Plane[0]
}
