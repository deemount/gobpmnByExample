package events

import (
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/events/elements"
)

type ProcessEventsElementsRepository interface {
	SetStartEvent(num int)
	GetStartEvent(num int) *elements.StartEvent
	SetEndEvent(num int)
	GetEndEvent(num int) END_EVENT_PTR
}
