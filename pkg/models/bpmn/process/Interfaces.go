package process

import (
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/events"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/flow"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/tasks"
)

// ProcessRepository ...
type ProcessRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	events.ProcessEventsElementsRepository
	tasks.TasksElementsRepository
	flow.FlowSequenceFlow

	SetIsExecutable(isExec bool)
	GetIsExecutable() *bool
}
