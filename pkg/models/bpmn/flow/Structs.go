package flow

import (
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/canvas"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
)

// flow
type DelegateParameter struct {
	SF    *SequenceFlow
	ED    *canvas.Edge
	WP    []canvas.Waypoint
	BS    canvas.Bounds  // bounds ref by value
	BSPTR *canvas.Bounds // bounds pointer to get bounds of last element
	ST    string         // source type
	TT    string         // target type
	T     string         // typ
	N     string         // name
	H     []string       // hash
}

type SourceTargetRef struct {
	SourceRef string `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef string `xml:"targetRef,attr" json:"targetRef,omitempty"`
}

// SequenceFlow ...
type SequenceFlow struct {
	impl.BaseAttributes
	SourceTargetRef
}

// TSequenceFlow ...
type TSequenceFlow struct {
	impl.BaseAttributes
	SourceTargetRef
}
