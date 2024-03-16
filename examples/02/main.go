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
		Def                  core.DefinitionsRepository
		IsExecutable         bool
		TenantProcess        gobpmn_hash.Injection
		TenantStartEvent     gobpmn_hash.Injection
		FromTenantStartEvent gobpmn_hash.Injection
		TenantTask           gobpmn_hash.Injection
		FromTenantTask       gobpmn_hash.Injection
		TenantEndEvent       gobpmn_hash.Injection
	}
)

/*
 * @Initiate
 */

// New ...
func New() Proxy {

	c := count.In(ExampleProcess{})
	p := hash.Inject(ExampleProcess{}).(ExampleProcess)

	//log.Printf("ExampleProcess: %+v\n", c)

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
	p.setTenantProcessArgs()
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

// elements ...
func (p *ExampleProcess) elements() {
	p.tenantProcess().SetStartEvent(1)
	p.tenantProcess().SetTask(1)
	p.tenantProcess().SetEndEvent(1)
	p.tenantProcess().SetSequenceFlow(2)
}

// setTenantProcess ...
func (p *ExampleProcess) setTenantProcessArgs() {
	p.tenantProcess().SetID("process", p.TenantProcess.Suffix)
	p.tenantProcess().SetIsExecutable(p.IsExecutable)
}

// setStartEvent ...
func (p *ExampleProcess) setStartEvent() {
	el := p.tenantStartEvent()
	el.SetID("startevent", p.TenantStartEvent.Suffix)
	el.SetName("Begin of Process")
	el.SetOutgoing(1)
	el.GetOutgoing(0).SetFlow(p.TenantTask.Suffix)
	p.fromStartEvent()
}

// fromStartEvent ...
func (p *ExampleProcess) fromStartEvent() {
	el := p.tenantProcess().GetSequenceFlow(0)
	el.SetID("flow", p.FromTenantStartEvent.Suffix)
	el.SetSourceRef("startevent", p.TenantStartEvent.Suffix)
	el.SetTargetRef("activity", p.TenantTask.Suffix)
}

// setTask ...
func (p *ExampleProcess) setTask() {
	el := p.tenantTask()
	el.SetID("activity", p.TenantTask.Suffix)
	el.SetName("Task")
	el.SetIncoming(1)
	el.GetIncoming(0).SetFlow(p.FromTenantStartEvent.Suffix)
	el.SetOutgoing(1)
	el.GetOutgoing(0).SetFlow(p.TenantEndEvent.Suffix)
	p.fromTask()
}

// fromTask ...
func (p *ExampleProcess) fromTask() {
	el := p.tenantProcess().GetSequenceFlow(1)
	el.SetID("flow", p.FromTenantTask.Suffix)
	el.SetSourceRef("activity", p.TenantTask.Suffix)
	el.SetTargetRef("endevent", p.TenantEndEvent.Suffix)
}

// setEndEvent ...
func (p *ExampleProcess) setEndEvent() {
	el := p.tenantProcess().GetEndEvent(0)
	el.SetID("endevent", p.TenantEndEvent.Suffix)
	el.SetName("End of Process")
	el.SetIncoming(1)
	el.GetIncoming(0).SetFlow(p.FromTenantTask.Suffix)
}

/**** Default Setter/Getter ****/

func (p *ExampleProcess) attributes() {
	p.Def.SetDefaultAttributes()
}

func (p ExampleProcess) tenantProcess() *process.Process {
	return p.Def.GetProcess(0)
}

func (p ExampleProcess) tenantStartEvent() *elements.StartEvent {
	return p.tenantProcess().GetStartEvent(0)
}

func (p ExampleProcess) tenantTask() *tasks.Task {
	return p.tenantProcess().GetTask(0)
}

/*
 * @Main
 */

// Main ...
func main() {

	exampleProcess := New().Build().Call()
	builder := gobpmn_builder.New()
	builder.SetDefinitionsByArg(exampleProcess)
	builder.Build()

}
