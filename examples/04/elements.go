package main

// elements ...
func (p *ExampleProcess) elements() {
	// Collaboration
	collaboration := p.Def.GetCollaboration()
	collaboration.SetID("collaboration", p.Collaboration.Suffix)
	collaboration.SetParticipant(2)
	collaboration.SetMessageFlow(2)
	// Processes
	// Support
	p.supportProcess().SetStartEvent(1)
	p.supportProcess().SetEndEvent(1)
	p.supportProcess().SetTask(2)
	p.supportProcess().SetSequenceFlow(3)
	// Customer
	p.customerProcess().SetStartEvent(1)
	p.customerProcess().SetEndEvent(1)
	p.customerProcess().SetTask(2)
	p.customerProcess().SetIntermediateCatchEvent(1)
	p.customerProcess().SetSequenceFlow(4)
}
