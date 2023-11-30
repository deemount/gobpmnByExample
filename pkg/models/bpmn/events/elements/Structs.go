package elements

import (
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/canvas"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/marker"
)

// DelegateParameter ...
type DelegateParameter struct {
	SE     *StartEvent
	EE     *EndEvent
	SH     *canvas.Shape
	BS     canvas.Bounds
	BSPTR  *canvas.Bounds
	WPPREV *canvas.Waypoint
	T      string
	N      string
	H      []string
}

// EndEvent ...
type EndEvent struct {
	impl.BaseAttributes
	Incoming []marker.Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
}

// TEndEvent ...
type TEndEvent struct {
	impl.BaseAttributes
	Incoming []marker.Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
}

// StartEvent ...
type StartEvent struct {
	impl.BaseAttributes
	Outgoing []marker.Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TStartEvent ...
type TStartEvent struct {
	impl.BaseAttributes
	Outgoing []marker.Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}
