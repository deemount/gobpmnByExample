package main

import (
	gobpmn_builder "github.com/deemount/gobpmnBuilder"
	gobpmn_counter "github.com/deemount/gobpmnCounter"
	gobpmn_hash "github.com/deemount/gobpmnHash"
	"github.com/deemount/gobpmnModels/pkg/core"
	"github.com/deemount/gobpmnModels/pkg/events/elements"
	"github.com/deemount/gobpmnModels/pkg/process"
	"github.com/deemount/gobpmnModels/pkg/tasks"
)

// Global variables ...
var (
	build gobpmn_builder.Builder
	count gobpmn_counter.Quantities
	hash  gobpmn_hash.Injection
)

// Structure ...
type (

	// Proxy ...
	Proxy interface {
		Build() ExampleProcess
	}

	// ExampleProcess ...
	ExampleProcess struct {
		Def            core.DefinitionsRepository
		IsExecutable   bool
		Process        gobpmn_hash.Injection
		StartEvent     gobpmn_hash.Injection
		FromStartEvent gobpmn_hash.Injection
		Task           gobpmn_hash.Injection
		FromTask       gobpmn_hash.Injection
		EndEvent       gobpmn_hash.Injection
	}
)

/*
 * @Initiate
 */

// New ...
func New() Proxy {

	c := count.In(ExampleProcess{})
	p := hash.Inject(ExampleProcess{}).(ExampleProcess)

	p.Def = core.NewDefinitions()
	p.Def.SetDefaultAttributes()

	build.Defaults(p.Def, c)

	return p
}

/*
 * @Global Methods
 */

// Build ...
func (p ExampleProcess) Build() ExampleProcess {
	p.elements()
	p.attributes()
	p.setProcess()
	p.setStartEvent()
	p.setTask()
	p.setEndEvent()
	return p
}

// Call ...
func (p ExampleProcess) Call() core.DefinitionsRepository {
	return p.Def
}

/*
 * @Local Methods
 */

// attributes ...
func (p *ExampleProcess) attributes() {
	p.Def.SetDefaultAttributes()
}

// elements ...
func (p *ExampleProcess) elements() {
	p.process().SetStartEvent(1)
	p.process().SetTask(1)
	p.process().SetEndEvent(1)
	p.process().SetSequenceFlow(2)
}

/****************************************************************************************/

func (p ExampleProcess) process() *process.Process {
	return p.Def.GetProcess(0)
}

func (p ExampleProcess) startEvent() *elements.StartEvent {
	return p.process().GetStartEvent(0)
}

func (p ExampleProcess) task() *tasks.Task {
	return p.process().GetTask(0)
}

func (p ExampleProcess) endEvent() *elements.EndEvent {
	return p.process().GetEndEvent(0)
}

/****************************************************************************************/

// setProcess ...
func (p *ExampleProcess) setProcess() {
	p.process().SetID("process", p.Process.Suffix)
	p.process().SetIsExecutable(p.IsExecutable)
}

// setStartEvent ...
func (p *ExampleProcess) setStartEvent() {
	el := p.startEvent()
	el.SetID("startevent", p.StartEvent.Suffix)
	el.SetName("Begin of Process")
	el.SetOutgoing(1)
	el.GetOutgoing(0).SetFlow(p.Task.Suffix)
	p.fromStartEvent()
}

// fromStartEvent ...
func (p *ExampleProcess) fromStartEvent() {
	el := p.process().GetSequenceFlow(0)
	el.SetID("flow", p.FromStartEvent.Suffix)
	el.SetSourceRef("startevent", p.StartEvent.Suffix)
	el.SetTargetRef("activity", p.Task.Suffix)
}

// setTask ...
func (p *ExampleProcess) setTask() {
	el := p.task()
	el.SetID("activity", p.Task.Suffix)
	el.SetName("Task")
	el.SetIncoming(1)
	el.GetIncoming(0).SetFlow(p.FromStartEvent.Suffix)
	el.SetOutgoing(1)
	el.GetOutgoing(0).SetFlow(p.EndEvent.Suffix)
	p.fromTask()
}

// fromTask ...
func (p *ExampleProcess) fromTask() {
	el := p.process().GetSequenceFlow(1)
	el.SetID("flow", p.FromTask.Suffix)
	el.SetSourceRef("activity", p.Task.Suffix)
	el.SetTargetRef("event", p.EndEvent.Suffix)
}

// setEndEvent ...
func (p *ExampleProcess) setEndEvent() {
	el := p.endEvent()
	el.SetID("event", p.EndEvent.Suffix)
	el.SetName("End of Process")
	el.SetIncoming(1)
	el.GetIncoming(0).SetFlow(p.FromTask.Suffix)
}

/*
 * @Main
 */

// Main ...
func main() {

	exampleProcess := New().Build().Call()
	builder := gobpmn_builder.New(gobpmn_builder.WithPath("files/bpmn2", "files/json2"), gobpmn_builder.WithCounter("files/bpmn2"))
	builder.SetDefinitionsByArg(exampleProcess)
	builder.Build()

}
