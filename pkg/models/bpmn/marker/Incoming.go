package marker

import "fmt"

// NewIncoming ...
func NewIncoming() IncomingRepository {
	return &Incoming{}
}

// SetFlow ...
func (incoming *Incoming) SetFlow(suffix interface{}) {
	incoming.Flow = fmt.Sprintf("Flow_%s", suffix)
}

// GetFlow ...
func (incoming Incoming) GetFlow() *string {
	return &incoming.Flow
}
