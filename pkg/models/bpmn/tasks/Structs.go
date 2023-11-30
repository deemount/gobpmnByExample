package tasks

import (
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/canvas"
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/flow"
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/marker"
)

// DelegateParameter ...
type DelegateParameter struct {
	TA     *Task
	SF     *flow.SequenceFlow
	SH     *canvas.Shape
	BS     canvas.Bounds
	WPPREV *canvas.Waypoint // second waypoint of previous edge
	T      string
	N      string
	H      []string
}

// Tasks ...
type Tasks struct {
	Task TASK_SLC `xml:"bpmn:task,omitempty" json:"task,omitempty"`
}

// TTasks ...
type TTasks struct {
	Task TASK_SLC `xml:"task,omitempty" json:"task,omitempty"`
}

// Task ...
type Task struct {
	impl.BaseAttributes
	marker.IncomingOutgoing
}

// TTask ...
type TTask struct {
	impl.BaseAttributes
	marker.TIncomingOutgoing
}
