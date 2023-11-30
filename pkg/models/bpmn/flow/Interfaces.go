package flow

import (
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
)

// FlowBaseReference ...
type FlowBaseReferences interface {
	SetSourceRef(typ string, sourceRef interface{})
	GetSourceRef() impl.STR_PTR
	SetTargetRef(typ string, targetRef interface{})
	GetTargetRef() impl.STR_PTR
}

// FlowSequenceFlow ...
type FlowSequenceFlow interface {
	SetSequenceFlow(num int)
	GetSequenceFlow(num int) *SequenceFlow
}

// FlowBase ...
type FlowBase interface {
	impl.IFBaseID
	impl.IFBaseName
}

// SequenceFlowRepository ...
type SequenceFlowRepository interface {
	FlowBase
	FlowBaseReferences
}
