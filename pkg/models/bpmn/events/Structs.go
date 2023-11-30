package events

import (
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/canvas"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/events/elements"
)

type DelegateParameter struct {
	SE     *elements.StartEvent
	EE     *elements.EndEvent
	SH     *canvas.Shape
	BS     canvas.Bounds
	WPPREV *canvas.Waypoint // previous waypoint
	T      string           // Typ
	N      string           // Name
	H      []string         // Hash
}

// ProcessEvents ...
type ProcessEvents struct {
	StartEvent START_EVENT_SLC `xml:"bpmn:startEvent,omitemnpty" json:"startEvent,omitempty" csv:"-"`
	EndEvent   END_EVENT_SLC   `xml:"bpmn:endEvent,omitempty" json:"endEvent,omitempty"`
}

// TProcessEvents ...
type TProcessEvents struct {
	StartEvent TSTART_EVENT_SLC `xml:"startEvent,omitemnpty" json:"startEvent,omitempty" csv:"-"`
	EndEvent   TEND_EVENT_SLC   `xml:"endEvent,omitempty" json:"endEvent,omitempty"`
}
