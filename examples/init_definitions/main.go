package main

import (
	"log"

	gobpmn_hash "github.com/deemount/gobpmnHash"
	"github.com/deemount/gobpmnModels/pkg/core"
	"github.com/deemount/gobpmnModels/pkg/process"
)

/*
 * @Proxy
 */

var hash gobpmn_hash.Injection

type Proxy interface {
	Build() ExampleProcess
}

/*
 * @ExampleProcess
 */

type (

	// ExampleProcess ...
	ExampleProcess struct {
		Def core.DefinitionsRepository
		Pool
		Tenant
	}

	// Pool
	Pool struct {
		TenantProcess gobpmn_hash.Injection
	}

	// Tenant
	Tenant struct {
		TenantStartEvent gobpmn_hash.Injection
	}
)

func New() Proxy {
	p := hash.Inject(ExampleProcess{}).(ExampleProcess)
	p.Def = core.NewDefinitions()
	hash.Create(p.Def)
	return &p
}

// Build sets the elements
func (p ExampleProcess) Build() ExampleProcess {
	p.tenantProcess().SetStartEvent(1)
	return p
}

// Def ...
func (p ExampleProcess) Call() core.DefinitionsRepository {
	return p.Def
}

// process returns the first process
func (p ExampleProcess) tenantProcess() *process.Process {
	return p.Def.GetProcess(0)
}

/*
 * @Main
 */

// Main ...
func main() {
	builder := New().Build().Call()
	log.Printf("main.go: %+v", builder)
}
