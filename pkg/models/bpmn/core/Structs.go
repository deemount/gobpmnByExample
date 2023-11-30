package core

import (
	"encoding/xml"

	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/canvas"
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/process"
)

// DefinitionsBaseElements ...
type DefinitionsBaseElements struct {
	Process process.PROCESS_SLC `xml:"bpmn:process,omitempty" json:"process"`
	Diagram []canvas.Diagram    `xml:"bpmndi:BPMNDiagram,omitempty" json:"diagram"`
}

// TDefinitionsBaseElements ...
type TDefinitionsBaseElements struct {
	Process process.TPROCESS_SLC `xml:"process,omitempty" json:"process"`
	Diagram []canvas.TDiagram    `xml:"BPMNDiagram,omitempty" json:"diagram"`
}

// Definitions represents the root element
type Definitions struct {
	XMLName         xml.Name `xml:"bpmn:definitions" json:"-"`
	Bpmn            string   `xml:"xmlns:bpmn,attr" json:"-"`
	BpmnDI          string   `xml:"xmlns:bpmndi,attr" json:"-"`
	DC              string   `xml:"xmlns:dc,attr,omitempty" json:"-"`
	ID              string   `xml:"id,attr" json:"id"`
	TargetNamespace string   `xml:"targetNamespace,attr" json:"-"`
	DefinitionsBaseElements
}

// TDefinitions ...
type TDefinitions struct {
	XMLName xml.Name `xml:"definitions" json:"-"`
	ID      string   `xml:"id,attr" json:"id"`
	TDefinitionsBaseElements
}
