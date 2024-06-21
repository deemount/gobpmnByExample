package main

import "github.com/deemount/gobpmnModels/pkg/core"

// New ...
func New() Proxy {

	quantities := count.In(ExampleProcess{})
	p := hash.Inject(ExampleProcess{}).(ExampleProcess)

	p.Def = core.NewDefinitions()
	p.Def.SetDefaultAttributes()

	build.Defaults(p.Def, quantities)

	return p
}

/*
 * @Global Methods
 */

// Build ...
func (p ExampleProcess) Build() ExampleProcess {
	p.Def.SetDefaultAttributes()
	p.elements()
	// Collaboration
	p.setParticipants()
	p.setMessageFlows()
	p.setSupportPool()
	p.setCustomerPool()
	// Customer Support
	p.setSupportProcess()
	p.setSupportStartEvent()
	p.setIncomingClaimTask()
	p.setDenyClaimTask()
	p.setSupportEndEvent()
	// Customer
	p.setCustomerProcess()
	p.setCustomerStartEvent()
	p.setNoticeTask()
	p.setWaitingTask()
	p.setRefusalTask()
	p.setCustomerEndEvent()
	return p
}

// Call ...
func (p ExampleProcess) Call() core.DefinitionsRepository {
	return p.Def
}
