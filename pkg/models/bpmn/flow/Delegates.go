package flow

import "github.com/deemount/gobpmnByExample/pkg/models/bpmn/canvas"

// SetSequenceFlow ...
func SetSequenceFlow(p DelegateParameter) {

	if p.T == "" {
		p.T = "flow"
	}

	p.SF.SetID(p.T, p.H[0])         // set id with first hash
	p.SF.SetSourceRef(p.ST, p.H[1]) // set source ref with second hash
	p.SF.SetTargetRef(p.TT, p.H[2]) // set target ref with third hash
	if p.ED != nil {
		canvas.SetEdge(
			canvas.DelegateParameter{
				E:     p.ED,
				T:     p.T,
				BSPTR: p.BSPTR, // Bounds Pointer (getting the element size before >>this<< edge)
				ST:    p.ST,
				TT:    p.TT,
				H:     p.H[0],
				W:     p.WP})
	}
}
