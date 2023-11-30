package process

import (
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/events"
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/flow"
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/tasks"
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
