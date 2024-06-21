package main

// setParticipants ...
func (p *ExampleProcess) setParticipants() {
	// Support
	support := p.Def.GetCollaboration().GetParticipant(0)
	support.SetID("participant", p.SupportID.Suffix)
	support.SetName("Support")
	support.SetProcessRef("process", p.SupportProcess.Suffix)
	// Customer
	customer := p.Def.GetCollaboration().GetParticipant(1)
	customer.SetID("participant", p.CustomerID.Suffix)
	customer.SetName("Customer")
	customer.SetProcessRef("process", p.CustomerProcess.Suffix)
}

// setMessageFlows ...
func (p *ExampleProcess) setMessageFlows() {
	// Claim
	messageClaim := p.Def.GetCollaboration().GetMessageFlow(0)
	messageClaim.SetID("flow", p.SupportToCustomerMessage.Suffix)
	messageClaim.SetName("Claim")
	messageClaim.SetSourceRef("activity", p.NoticeTask.Suffix)
	messageClaim.SetTargetRef("activity", p.DenyClaimTask.Suffix)
	// Refusal
	messageRefusal := p.Def.GetCollaboration().GetMessageFlow(1)
	messageRefusal.SetID("flow", p.CustomerToSupportMessage.Suffix)
	messageRefusal.SetName("Refusal")
	messageRefusal.SetSourceRef("activity", p.DenyClaimTask.Suffix)
	messageRefusal.SetTargetRef("activity", p.RefusalTask.Suffix)
}
