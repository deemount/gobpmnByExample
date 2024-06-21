package main

import (
	"github.com/deemount/gobpmnModels/pkg/events/elements"
	"github.com/deemount/gobpmnModels/pkg/process"
	"github.com/deemount/gobpmnModels/pkg/tasks"
)

func (p ExampleProcess) customerProcess() *process.Process {
	return p.Def.GetProcess(1)
}

func (p ExampleProcess) customerStartEvent() *elements.StartEvent {
	return p.customerProcess().GetStartEvent(0)
}

func (p ExampleProcess) noticeTask() *tasks.Task {
	return p.customerProcess().GetTask(0)
}

func (p ExampleProcess) waitingTask() *elements.IntermediateCatchEvent {
	return p.customerProcess().GetIntermediateCatchEvent(0)
}

func (p ExampleProcess) refusalTask() *tasks.Task {
	return p.customerProcess().GetTask(1)
}

func (p ExampleProcess) customerEndEvent() *elements.EndEvent {
	return p.customerProcess().GetEndEvent(0)
}

/****************************************************************************************/

// setCustomerPool ...
func (p *ExampleProcess) setCustomerPool() {
}

// setCustomerProcess ...
func (p *ExampleProcess) setCustomerProcess() {
	p.customerProcess().SetID("process", p.CustomerProcess.Suffix)
	p.customerProcess().SetIsExecutable(p.CustomerIsExecutable)
}

// setCustomerStartEvent ...
func (p *ExampleProcess) setCustomerStartEvent() {
	el := p.customerStartEvent()
	el.SetID("startevent", p.CustomerStartEvent.Suffix)
	el.SetName("Begin of Process")
	el.SetOutgoing(1)
	el.GetOutgoing(0).SetFlow(p.NoticeTask.Suffix)
	p.fromCustomerStartEvent()
}

// fromCustomerStartEvent ...
func (p *ExampleProcess) fromCustomerStartEvent() {
	el := p.customerProcess().GetSequenceFlow(0)
	el.SetID("flow", p.FromCustomerStartEvent.Suffix)
	el.SetSourceRef("startevent", p.CustomerStartEvent.Suffix)
	el.SetTargetRef("activity", p.NoticeTask.Suffix)
}

// setTask ...
func (p *ExampleProcess) setNoticeTask() {
	el := p.noticeTask()
	el.SetID("activity", p.NoticeTask.Suffix)
	el.SetName("Notice")
	el.SetIncoming(1)
	el.GetIncoming(0).SetFlow(p.FromCustomerStartEvent.Suffix)
	el.SetOutgoing(1)
	el.GetOutgoing(0).SetFlow(p.WaitingTask.Suffix)
	p.fromNoticeTask()
}

// fromNoticeTask ...
func (p *ExampleProcess) fromNoticeTask() {
	el := p.customerProcess().GetSequenceFlow(1)
	el.SetID("flow", p.FromNoticeTask.Suffix)
	el.SetSourceRef("activity", p.NoticeTask.Suffix)
	el.SetTargetRef("activity", p.WaitingTask.Suffix)
}

// setTask ...
func (p *ExampleProcess) setWaitingTask() {
	el := p.waitingTask()
	el.SetID("activity", p.WaitingTask.Suffix)
	el.SetName("Waiting")
	el.SetIncoming(1)
	el.GetIncoming(0).SetFlow(p.FromNoticeTask.Suffix)
	el.SetOutgoing(1)
	el.GetOutgoing(0).SetFlow(p.RefusalTask.Suffix)
	p.fromWaitingTask()
}

// fromWaitingTask ...
func (p *ExampleProcess) fromWaitingTask() {
	el := p.customerProcess().GetSequenceFlow(2)
	el.SetID("flow", p.FromWaitingTask.Suffix)
	el.SetSourceRef("activity", p.WaitingTask.Suffix)
	el.SetTargetRef("activity", p.RefusalTask.Suffix)
}

// setTask ...
func (p *ExampleProcess) setRefusalTask() {
	el := p.refusalTask()
	el.SetID("activity", p.RefusalTask.Suffix)
	el.SetName("Refusal")
	el.SetIncoming(1)
	el.GetIncoming(0).SetFlow(p.FromWaitingTask.Suffix)
	el.SetOutgoing(1)
	el.GetOutgoing(0).SetFlow(p.CustomerEndEvent.Suffix)
	p.fromRefusalTask()
}

// fromRefusalTask ...
func (p *ExampleProcess) fromRefusalTask() {
	el := p.customerProcess().GetSequenceFlow(3)
	el.SetID("flow", p.FromRefusalTask.Suffix)
	el.SetSourceRef("activity", p.RefusalTask.Suffix)
	el.SetTargetRef("activity", p.CustomerEndEvent.Suffix)
}

// setCustomerEndEvent ...
func (p *ExampleProcess) setCustomerEndEvent() {
	el := p.customerEndEvent()
	el.SetID("event", p.CustomerEndEvent.Suffix)
	el.SetName("End of Process")
	el.SetIncoming(1)
	el.GetIncoming(0).SetFlow(p.FromRefusalTask.Suffix)
}
