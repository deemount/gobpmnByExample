package main

import (
	"github.com/deemount/gobpmnModels/pkg/events/elements"
	"github.com/deemount/gobpmnModels/pkg/process"
	"github.com/deemount/gobpmnModels/pkg/tasks"
)

func (p ExampleProcess) supportProcess() *process.Process {
	return p.Def.GetProcess(0)
}

func (p ExampleProcess) supportStartEvent() *elements.StartEvent {
	return p.supportProcess().GetStartEvent(0)
}

func (p ExampleProcess) incomingClaimTask() *tasks.Task {
	return p.supportProcess().GetTask(0)
}

func (p ExampleProcess) denyClaimTask() *tasks.Task {
	return p.supportProcess().GetTask(1)
}

func (p ExampleProcess) supportEndEvent() *elements.EndEvent {
	return p.supportProcess().GetEndEvent(0)
}

/****************************************************************************************/

// setSupportPool ... ...
func (p *ExampleProcess) setSupportPool() {
}

// setSupportProcess ...
func (p *ExampleProcess) setSupportProcess() {
	p.supportProcess().SetID("process", p.SupportProcess.Suffix)
	p.supportProcess().SetIsExecutable(p.SupportIsExecutable)
}

// setSupportStartEvent ...
func (p *ExampleProcess) setSupportStartEvent() {
	el := p.supportStartEvent()
	el.SetID("startevent", p.SupportStartEvent.Suffix)
	el.SetName("Begin of Process")
	el.SetOutgoing(1)
	el.GetOutgoing(0).SetFlow(p.IncomingClaimTask.Suffix)
	p.fromSupportStartEvent()
}

// fromSupportStartEvent ...
func (p *ExampleProcess) fromSupportStartEvent() {
	el := p.supportProcess().GetSequenceFlow(0)
	el.SetID("flow", p.FromSupportStartEvent.Suffix)
	el.SetSourceRef("startevent", p.SupportStartEvent.Suffix)
	el.SetTargetRef("activity", p.IncomingClaimTask.Suffix)
}

// setTask ...
func (p *ExampleProcess) setIncomingClaimTask() {
	el := p.incomingClaimTask()
	el.SetID("activity", p.IncomingClaimTask.Suffix)
	el.SetName("Incoming Claim")
	el.SetIncoming(1)
	el.GetIncoming(0).SetFlow(p.FromSupportStartEvent.Suffix)
	el.SetOutgoing(1)
	el.GetOutgoing(0).SetFlow(p.DenyClaimTask.Suffix)
	p.fromIncomingClaimTask()
}

// fromIncomingClaimTask ...
func (p *ExampleProcess) fromIncomingClaimTask() {
	el := p.supportProcess().GetSequenceFlow(1)
	el.SetID("flow", p.FromIncomingClaimTask.Suffix)
	el.SetSourceRef("activity", p.IncomingClaimTask.Suffix)
	el.SetTargetRef("activity", p.SupportEndEvent.Suffix)
}

// setTask ...
func (p *ExampleProcess) setDenyClaimTask() {
	el := p.denyClaimTask()
	el.SetID("activity", p.DenyClaimTask.Suffix)
	el.SetName("Deny Claim")
	el.SetIncoming(1)
	el.GetIncoming(0).SetFlow(p.FromIncomingClaimTask.Suffix)
	el.SetOutgoing(1)
	el.GetOutgoing(0).SetFlow(p.SupportEndEvent.Suffix)
	p.fromDenyClaimTask()
}

// fromDenyClaimTask ...
func (p *ExampleProcess) fromDenyClaimTask() {
	el := p.supportProcess().GetSequenceFlow(2)
	el.SetID("flow", p.FromDenyClaimTask.Suffix)
	el.SetSourceRef("activity", p.DenyClaimTask.Suffix)
	el.SetTargetRef("activity", p.SupportEndEvent.Suffix)
}

// setSupportEndEvent ...
func (p *ExampleProcess) setSupportEndEvent() {
	el := p.supportEndEvent()
	el.SetID("event", p.SupportEndEvent.Suffix)
	el.SetName("End of Process")
	el.SetIncoming(1)
	el.GetIncoming(0).SetFlow(p.FromDenyClaimTask.Suffix)
}
