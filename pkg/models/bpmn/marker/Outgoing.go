package marker

import "fmt"

// NewOutgoing ...
func NewOutgoing() OutgoingRepository {
	return &Outgoing{}
}

// SetFlow ...
func (outgoing *Outgoing) SetFlow(suffix interface{}) {
	outgoing.Flow = fmt.Sprintf("Flow_%s", suffix)
}

// GetFlow ...
func (outgoing Outgoing) GetFlow() *string {
	return &outgoing.Flow
}
