package flow

import (
	"fmt"

	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
)

// NewSequenceFlow ...
func NewSequenceFlow() SequenceFlowRepository {
	return &SequenceFlow{}
}

// SetID ...
func (sequenceFlow *SequenceFlow) SetID(typ string, suffix interface{}) {
	switch typ {
	case "flow":
		sequenceFlow.ID = fmt.Sprintf("Flow_%v", suffix)
		break
	}
}

// SetName ...
func (sequenceFlow *SequenceFlow) SetName(name string) {
	sequenceFlow.Name = name
}

// SetSourceRef ...
func (sequenceFlow *SequenceFlow) SetSourceRef(typ string, sourceRef interface{}) {
	switch typ {
	case "activity":
		sequenceFlow.SourceRef = fmt.Sprintf("Activity_%s", sourceRef)
		break
	case "event":
		sequenceFlow.SourceRef = fmt.Sprintf("Event_%s", sourceRef)
		break
	case "startevent":
		sequenceFlow.SourceRef = fmt.Sprintf("StartEvent_%v", sourceRef)
		break
	}
}

// SetTargetRef ...
func (sequenceFlow *SequenceFlow) SetTargetRef(typ string, targetRef interface{}) {
	switch typ {
	case "activity":
		sequenceFlow.TargetRef = fmt.Sprintf("Activity_%s", targetRef)
		break
	case "event":
		sequenceFlow.TargetRef = fmt.Sprintf("Event_%s", targetRef)
		break
	case "startevent":
		sequenceFlow.TargetRef = fmt.Sprintf("StartEvent_%s", targetRef)
		break
	}
}

// GetID ...
func (sequenceFlow SequenceFlow) GetID() impl.STR_PTR {
	return &sequenceFlow.ID
}

// GetName ...
func (sequenceFlow SequenceFlow) GetName() impl.STR_PTR {
	return &sequenceFlow.Name
}

// GetSourceRef ...
func (sequenceFlow SequenceFlow) GetSourceRef() impl.STR_PTR {
	return &sequenceFlow.SourceRef
}

// GetTargetRef ...
func (sequenceFlow SequenceFlow) GetTargetRef() impl.STR_PTR {
	return &sequenceFlow.TargetRef
}
