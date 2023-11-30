package canvas

import "github.com/deemount/gobpmnByExamples/pkg/models/bpmn/impl"

type CanvasBoundsElements interface {
	SetBounds()
	GetBounds() *Bounds
}

// BoundsRepository ...
type BoundsRepository interface {
	impl.IFBaseCoordinates
	SetSize(width, height int)
	GetSize() (impl.INT_PTR, impl.INT_PTR)
	SetWidth(width int)
	GetWidth() impl.INT_PTR
	SetHeight(height int)
	GetHeight() impl.INT_PTR
}

// WaypointRepository ...
type WaypointRepository interface {
	impl.IFBaseCoordinates
}

// Diagram ...
type DiagramRepository interface {
	impl.IFBaseID
	SetPlane()
	GetPlane() *Plane
}

// Edge ...
type EdgeRepository interface {
	impl.IFBaseID
	impl.IFBaseElement
	SetWaypoint()
	SetWaypoints(num int)
	GetWaypoint(num int) *Waypoint
}

// PlaneRepository ...
type PlaneRepository interface {
	impl.IFBaseID
	impl.IFBaseElement
	SetAttrProcessElement(suffix string)
	SetAttrCollaborationElement(suffix string)
	SetShape(num int)
	GetShape(num int) *Shape
	SetEdge(num int)
	GetEdge(num int) *Edge
}

// ShapeRepository ...
type ShapeRepository interface {
	impl.IFBaseID
	impl.IFBaseElement
	CanvasBoundsElements
}
