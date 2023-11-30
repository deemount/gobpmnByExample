package marker

// marker
type DelegateParameter struct {
}

type IncomingOutgoing struct {
	Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

type TIncomingOutgoing struct {
	Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

// Incoming ...
type Incoming struct {
	Flow string `xml:",innerxml" json:"flow"`
}

// Outgoing ...
type Outgoing struct {
	Flow string `xml:",innerxml" json:"flow"`
}
