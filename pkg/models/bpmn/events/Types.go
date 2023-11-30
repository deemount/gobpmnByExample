package events

import "github.com/deemount/gobpmnByExamples/pkg/models/bpmn/events/elements"

type END_EVENT_PTR *elements.EndEvent
type START_EVENT_PTR *elements.StartEvent

type PROCESS_EVENTS_SLC []ProcessEvents

type END_EVENT_SLC []elements.EndEvent
type START_EVENT_SLC []elements.StartEvent

type TEND_EVENT_SLC []elements.TEndEvent
type TSTART_EVENT_SLC []elements.TStartEvent
