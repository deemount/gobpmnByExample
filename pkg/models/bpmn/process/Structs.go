package process

import (
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/events"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/flow"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/tasks"
)

// Process ...
type Process struct {
	impl.BaseAttributes
	events.ProcessEvents
	tasks.Tasks
	IsExecutable bool                `xml:"isExecutable,attr" json:"isExecutable,omitempty"`
	SequenceFlow []flow.SequenceFlow `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

// TProcess ...
type TProcess struct {
	impl.BaseAttributes
	events.TProcessEvents
	tasks.TTasks
	IsExecutable bool                 `xml:"isExecutable,attr" json:"isExecutable,omitempty"`
	SequenceFlow []flow.TSequenceFlow `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}
