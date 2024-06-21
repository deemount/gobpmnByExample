package main

import (
	gobpmn_hash "github.com/deemount/gobpmnHash"
	"github.com/deemount/gobpmnModels/pkg/core"
)

// Structure ...
type (

	// ExampleProcess ...
	ExampleProcess struct {
		Def          core.DefinitionsRepository
		IsExecutable bool
		Pool
		Message
		Support
		Customer
	}

	// Pool is a struct for the pool related elements and it's attributes and methods ...
	Pool struct {
		// configuration
		SupportIsExecutable  bool
		CustomerIsExecutable bool
		// pool related
		Collaboration   gobpmn_hash.Injection
		SupportID       gobpmn_hash.Injection
		SupportProcess  gobpmn_hash.Injection
		CustomerID      gobpmn_hash.Injection
		CustomerProcess gobpmn_hash.Injection
	}

	// Support ...
	Support struct {
		SupportStartEvent     gobpmn_hash.Injection
		FromSupportStartEvent gobpmn_hash.Injection
		IncomingClaimTask     gobpmn_hash.Injection
		FromIncomingClaimTask gobpmn_hash.Injection
		DenyClaimTask         gobpmn_hash.Injection
		FromDenyClaimTask     gobpmn_hash.Injection
		SupportEndEvent       gobpmn_hash.Injection
	}

	// Customer ...
	Customer struct {
		CustomerStartEvent          gobpmn_hash.Injection
		FromCustomerStartEvent      gobpmn_hash.Injection
		NoticeTask                  gobpmn_hash.Injection
		FromNoticeTask              gobpmn_hash.Injection
		WaitingTask                 gobpmn_hash.Injection
		TimerEventDefinitionWaiting gobpmn_hash.Injection
		FromWaitingTask             gobpmn_hash.Injection
		RefusalTask                 gobpmn_hash.Injection
		FromRefusalTask             gobpmn_hash.Injection
		CustomerEndEvent            gobpmn_hash.Injection
	}

	// Message ...
	Message struct {
		CustomerToSupportMessage gobpmn_hash.Injection
		SupportToCustomerMessage gobpmn_hash.Injection
	}
)
