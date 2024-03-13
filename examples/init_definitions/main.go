package main

import (
	"log"

	gobpmn_builder "github.com/deemount/gobpmnBuilder"
	gobpmn_counter "github.com/deemount/gobpmnCounter"
	gobpmn_hash "github.com/deemount/gobpmnHash"
	"github.com/deemount/gobpmnModels/pkg/core"
	"github.com/deemount/gobpmnModels/pkg/events/elements"
	"github.com/deemount/gobpmnModels/pkg/process"
)

/*
 * @ExampleProcess
 */

var build gobpmn_builder.Builder
var count gobpmn_counter.Quantities
var hash gobpmn_hash.Injection

type (

	// Proxy ...
	Proxy interface {
		Build() ExampleProcess
	}

	// ExampleProcess ...
	ExampleProcess struct {
		Def core.DefinitionsRepository
		Pool
		Tenant
	}

	// Pool
	Pool struct {
		TenantIsExecutable bool
		TenantProcess      gobpmn_hash.Injection
	}

	// Tenant
	Tenant struct {
		TenantStartEvent gobpmn_hash.Injection
	}
)

func New() Proxy {

	c := count.In(ExampleProcess{})
	log.Printf("main.go: %+v", c)

	p := hash.Inject(ExampleProcess{}).(ExampleProcess)
	log.Printf("main.go: %+v", p)

	p.Def = core.NewDefinitions()
	p.Def.SetDefaultAttributes()

	build.Defaults(p.Def, c)

	return &p
}

// Build sets the elements
func (p ExampleProcess) Build() ExampleProcess {
	p.setTenantProcessArgs()
	p.tenantProcess().SetStartEvent(1)
	p.setTenantStartEventID()
	return p
}

// Call returns the definitions reference ...
func (p ExampleProcess) Call() core.DefinitionsRepository {
	return p.Def
}

// setProcess ...
func (p *ExampleProcess) setTenantProcessArgs() {
	p.tenantProcess().SetID("process", p.TenantProcess.Suffix)
	p.tenantProcess().SetIsExecutable(p.TenantIsExecutable)
}

func (p *ExampleProcess) setTenantStartEventID() {
	p.tenantStartEvent().SetID("startevent", p.TenantStartEvent.Suffix)
}

func (p ExampleProcess) tenantProcess() *process.Process {
	return p.Def.GetProcess(0)
}

func (p ExampleProcess) tenantStartEvent() *elements.StartEvent {
	return p.tenantProcess().GetStartEvent(0)
}

/*
 * @Main
 */

// Main ...
func main() {

	exampleProcess := New().Build().Call()
	log.Printf("main.go: %+v", exampleProcess)

	builder := gobpmn_builder.New()
	builder.SetDefinitionsByArg(exampleProcess)
	builder.Build()

}
