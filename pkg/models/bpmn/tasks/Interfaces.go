package tasks

import (
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/marker"
)

type TasksBaseAttributes interface {
	impl.IFBaseID
	impl.IFBaseName
}

type TasksMarkers interface {
	SetIncoming(num int)
	GetIncoming(num int) *marker.Incoming
	SetOutgoing(num int)
	GetOutgoing(num int) *marker.Outgoing
}

type TasksBase interface {
	TasksBaseAttributes
	TasksMarkers
}

// TasksElementsRepository ...
type TasksElementsRepository interface {
	SetTask(num int)
	GetTask(num int) TASK_PTR
}

// TaskRepository ...
type TaskRepository interface {
	TasksBase
	String() string
}
