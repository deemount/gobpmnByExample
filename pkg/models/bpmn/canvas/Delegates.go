package canvas

import (
	"fmt"
	"log"
	"strconv"
)

// Center ...
type Center struct {
	x, y int
}

func center(bounds Bounds) Center {
	return Center{
		x: bounds.X + (bounds.Width / 2),
		y: bounds.Y + (bounds.Height / 2),
	}
}

// Big Delta calculates the difference between two elements
type Delta struct {
	x, y float64
}

func delta(a, b Point) Delta {
	return Delta{
		x: a.X - b.X,
		y: a.X - b.Y,
	}
}

// defaultDistanceTop ...
func (p DelegateParameter) defaultDistanceTop() int {
	return 50
}

// defaultDistanceLeft ...
func (p DelegateParameter) defaultDistanceLeft() int {
	return 50
}

// defaultEdgeLength ...
func (p DelegateParameter) defaultEdgeLength() int {
	return 55
}

// defaultCoordinates ...
func (p DelegateParameter) defaultCoordinates() (int, int) {
	var x, y int
	switch p.T {
	case "event":
		x, y = 179, 159
	case "startevent":
		x, y = 179, 159
	}
	return x, y
}

// defaultElementSize ...
func (p DelegateParameter) defaultElementSize() (int, int) {

	var width, height int
	var typ string

	if p.T == "flow" && p.ST != "" {
		typ = p.ST // the type, which is assigned the flow to
	} else {
		typ = p.T // the actually type
	}

	switch typ {
	case "activity":
		width, height = 100, 80
	case "event":
		width, height = 36, 36
	case "startevent":
		width, height = 36, 36
	}
	return width, height
}

// findCoordinatesByPreviousWaypoint ...
func (p *DelegateParameter) findCoordinatesByPreviousWaypoint() (int, int) {

	_, height := p.defaultElementSize()

	x := p.WPPREV.X
	y := p.WPPREV.Y - (height / 2)

	return x, y
}

// findCoordinatesByPreviousBounds ...
func (p *DelegateParameter) findCoordinatesByPreviousBounds() (int, int) {

	_, height := p.defaultElementSize()

	x := p.BSPTR.X + p.defaultDistanceLeft()
	y := (p.BSPTR.Y + (p.BSPTR.Height / 2)) - (height / 2)

	return x, y
}

// setBounds sets the map for the shape,
// sets default or given coordinates for the shape
// and sets the elements default or given size of the shape
func (p *DelegateParameter) setBounds() {

	p.S.SetBounds()

	// if coordinates of x and y are zero, decide between two conditions:
	// - if previous waypoint is not nil, find the coordinates by waypoint
	// -- if previous bounds is not nil, find the coordinates by bounds
	// -- else set default coordinates, when previous waypoint is nil
	if p.B.X == 0 && p.B.Y == 0 {

		if p.WPPREV != nil {
			p.B.X, p.B.Y = p.findCoordinatesByPreviousWaypoint()
		} else {
			if p.BSPTR != nil {
				p.B.X, p.B.Y = p.findCoordinatesByPreviousBounds()
			} else {
				p.B.X, p.B.Y = p.defaultCoordinates()
			}
		}

	}
	p.S.GetBounds().SetCoordinates(p.B.X, p.B.Y)

	// if width and height are zero, set default element size
	if p.B.Width == 0 && p.B.Height == 0 {
		p.B.Width, p.B.Height = p.defaultElementSize()
	}
	p.S.GetBounds().SetSize(p.B.Width, p.B.Height)

}

// setWaypoints ...
func (p *DelegateParameter) setWaypoints() {

	if len(p.W) > 2 {
		// many waypoints
		p.E.SetWaypoints(len(p.W))
	} else {
		// couple of waypoint
		p.E.SetWaypoint()
	}

	// waypoint sizes aren't given ...
	if len(p.W) == 0 {

		// calculate the first waypoint X and Y
		width, height := p.defaultElementSize()
		x1 := width + p.BSPTR.X
		y1 := (height / 2) + p.BSPTR.Y
		p.E.GetWaypoint(0).SetCoordinates(x1, y1)

		// calculate the second waypoint X and Y
		x2 := p.defaultEdgeLength() + x1
		y2 := y1
		p.E.GetWaypoint(1).SetCoordinates(x2, y2)

	} else {

		p.E.GetWaypoint(0).SetCoordinates(p.W[0].X, p.W[0].Y)
		p.E.GetWaypoint(1).SetCoordinates(p.W[1].X, p.W[1].Y)

		if len(p.W) > 2 {
			for i := 2; i < len(p.W); i++ {
				p.E.GetWaypoint(i).SetCoordinates(p.W[i].X, p.W[i].Y)
			}
		}

	}

}

/*
 *
 */

// SetShape ...
func SetShape(p DelegateParameter) {

	if p.S == nil && p.E == nil {
		log.Fatal("fatal: missing element pointer in models/bpmn/canvas/Delegates.go SetShape()")
	}

	// detect startevent
	if p.T == "startevent" {

		// if startevent selected, then count up the _BPMNShape_StartEvent ID
		newHash, _ := strconv.Atoi(p.H)
		newHash += 1
		p.S.SetID(p.T, fmt.Sprint(newHash))

	} else {
		p.S.SetID(p.T, p.H)
	}

	p.S.SetElement(p.T, p.H)
	p.setBounds()

}

// SetEdge ...
func SetEdge(p DelegateParameter) {
	p.E.SetID(p.T, p.H)
	p.E.SetElement(p.T, p.H)
	p.setWaypoints()
}
