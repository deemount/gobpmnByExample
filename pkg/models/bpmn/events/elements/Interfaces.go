package elements

import (
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/marker"
)

// EndEventRepository ...
type EndEventRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	marker.MarkerIncoming
	String() string
}

// TEndEventRepository ...
type TEndEventRepository interface {
	String() string
}

// StartEventRepository ...
type StartEventRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	marker.MarkerOutgoing
	String() string
}

// TStartEventRepository ...
type TStartEventRepository interface {
	String() string
}
