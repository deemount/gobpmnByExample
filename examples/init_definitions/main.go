package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	gobpmn_count "github.com/deemount/gobpmnCounter"
	gobpmn_hash "github.com/deemount/gobpmnHash"
	"github.com/deemount/gobpmnModels/pkg/core"
	"github.com/deemount/gobpmnModels/pkg/process"
)

/*
 * @ExampleProcess
 */

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

	var count gobpmn_count.Quantities

	c := count.In(ExampleProcess{})
	log.Printf("main.go: %+v", c)

	p := hash.Inject(ExampleProcess{}).(ExampleProcess)
	log.Printf("main.go: %+v", p)

	p.Def = core.NewDefinitions()
	hash.Create(p.Def)

	return &p
}

// Build sets the elements
func (p ExampleProcess) Build() ExampleProcess {
	p.setProcess()
	p.tenantProcess().SetStartEvent(1)
	return p
}

// Call returns the definitions reference ...
func (p ExampleProcess) Call() core.DefinitionsRepository {
	return p.Def
}

// setProcess ...
func (p *ExampleProcess) setProcess() {
	p.tenantProcess().SetID("process", p.TenantProcess.Suffix)
	p.tenantProcess().SetIsExecutable(p.TenantIsExecutable)
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

	var err error

	builder := New().Build().Call()
	log.Printf("main.go: %+v", builder)

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(&builder, " ", "  ")

	// create .bpmn file
	f, err := os.Create("../../files/bpmn/test.bpmn")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// add xml header
	w := []byte(fmt.Sprintf("%v", xml.Header+string(b)))

	// write bytes to file
	_, err = f.Write(w)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
