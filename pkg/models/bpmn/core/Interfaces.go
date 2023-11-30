package core

import (
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/canvas"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/process"
)

// DefinitionsElements ...
type DefinitionsElements interface {
	SetProcess(num int)
	GetProcess(num int) process.PROCESS_PTR
	SetDiagram(num int)
	GetDiagram(num int) canvas.DIAGRAM_PTR
}

// DefinitionsRepository ...
type DefinitionsRepository interface {
	impl.IFBaseID
	DefinitionsElements
	SetBpmn()
	SetBpmnDI()
	SetDC()
	SetTargetNamespace()
	SetMainElements(num int)
	SetDefaultAttributes()
}
