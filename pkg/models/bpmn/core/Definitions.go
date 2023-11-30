package core

import (
	"fmt"

	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/canvas"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"
	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/process"
	"github.com/deemount/gobpmnByExamples/pkg/utils"
)

var (
	schemaBpmnModel    = "http://www.omg.org/spec/BPMN/20100524/MODEL"
	schemaBpmnDI       = "http://www.omg.org/spec/BPMN/20100524/DI"
	schemaOMGDC        = "http://www.omg.org/spec/DD/20100524/DC"
	schemaBpmnIOSchema = "http://bpmn.io/schema/bpmn"
)

// NewDefinitions ...
func NewDefinitions() DefinitionsRepository {
	return &Definitions{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetBpmn ...
func (definitions *Definitions) SetBpmn() {
	definitions.Bpmn = schemaBpmnModel
}

// SetBpmnDI ...
func (definitions *Definitions) SetBpmnDI() {
	definitions.BpmnDI = schemaBpmnDI
}

// SetDC ...
func (definitions *Definitions) SetDC() {
	definitions.DC = schemaOMGDC
}

// SetID ...
func (definitions *Definitions) SetID(typ string, suffix interface{}) {
	definitions.ID = fmt.Sprintf("Definitions_%v", suffix)
}

// SetTargetNamespace ...
func (definitions *Definitions) SetTargetNamespace() {
	definitions.TargetNamespace = schemaBpmnIOSchema
}

/*** Make Elements ***/

/** BPMN **/

// SetProcess ...
func (definitions *Definitions) SetProcess(num int) {
	if num == 0 {
		num = 1
	}
	definitions.Process = make(process.PROCESS_SLC, num)
}

/** BPMNDI **/

// SetDiagram ...
func (definitions *Definitions) SetDiagram(num int) {
	definitions.Diagram = make(canvas.DIAGRAM_SLC, num)
}

/*
 * Default Settings
 */

// SetDefinitionsAttributes ...
func (definitions *Definitions) SetDefaultAttributes() {
	definitionsHash := utils.GenerateHash()
	definitions.SetBpmn()
	definitions.SetBpmnDI()
	definitions.SetDC()
	definitions.SetID("definitions", definitionsHash)
	definitions.SetTargetNamespace()
}

// SetMainElements ...
func (definitions *Definitions) SetMainElements(num int) {
	definitions.SetProcess(num)
	definitions.SetDiagram(1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (definitions Definitions) GetID() impl.STR_PTR {
	return &definitions.ID
}

/* Elements */

/** BPMN **/

// GetProcess ...
func (definitions Definitions) GetProcess(num int) process.PROCESS_PTR {
	return &definitions.Process[num]
}

/** BPMNDI **/

// SetDiagram ...
func (definitions Definitions) GetDiagram(num int) canvas.DIAGRAM_PTR {
	return &definitions.Diagram[num]
}

/*
 * Default String
 */

// String ...
func (definitions Definitions) String() string {
	return fmt.Sprintf("id=%v", definitions.ID)
}

// String ...
func (definitions TDefinitions) String() string {
	return fmt.Sprintf("id=%v", definitions.ID)
}
