package canvas

// canvas
// Shape: d *Shape, typ string, hash string, b Bounds
// Edge: d *Edge, typ string, hash string, w []Waypoint
type DelegateParameter struct {
	S      *Shape
	E      *Edge
	T      string // typ
	H      string // hash
	I      bool
	W      []Waypoint
	WPPREV *Waypoint // previous waypoint (needed in activity or event)
	B      Bounds
	BSPTR  *Bounds // bounds pointer (needed in flow)
	ST     string  // source ref
	TT     string  // target ref
	SZ     []int   // sizes
}

// Bounds ...
type Bounds struct {
	X      int `xml:"x,attr,omitempty" json:"x,omitempty"`
	Y      int `xml:"y,attr,omitempty" json:"y,omitempty"`
	Width  int `xml:"width,attr,omitempty"`
	Height int `xml:"height,attr,omitempty"`
}

// Waypoint ...
type Waypoint struct {
	X int `xml:"x,attr" json:"x,omitempty"`
	Y int `xml:"y,attr" json:"y,omitempty"`
}

// Diagram ...
type Diagram struct {
	ID    string  `xml:"id,attr" json:"id,omitempty"`
	Plane []Plane `xml:"bpmndi:BPMNPlane,omitempty" json:"plane,omitempty"`
}

// TDiagram ...
type TDiagram struct {
	ID    string   `xml:"id,attr" json:"id,omitempty"`
	Plane []TPlane `xml:"BPMNPlane,omitempty" json:"plane,omitempty"`
}

// Edge ...
type Edge struct {
	ID       string     `xml:"id,attr" json:"-"`
	Element  string     `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Waypoint []Waypoint `xml:"di:waypoint" json:"waypoint,omitempty"`
}

// TEdge ...
type TEdge struct {
	ID       string     `xml:"id,attr" json:"-"`
	Element  string     `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Waypoint []Waypoint `xml:"waypoint" json:"waypoint,omitempty"`
}

// Plane ...
type Plane struct {
	ID      string  `xml:"id,attr" json:"id,omitempty"`
	Element string  `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Shape   []Shape `xml:"bpmndi:BPMNShape" json:"shape,omitempty"`
	Edge    []Edge  `xml:"bpmndi:BPMNEdge" json:"edge,omitempty"`
}

// TPlane ...
type TPlane struct {
	ID      string   `xml:"id,attr" json:"id,omitempty"`
	Element string   `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Shape   []TShape `xml:"BPMNShape" json:"shape,omitempty"`
	Edge    []TEdge  `xml:"BPMNEdge" json:"edge,omitempty"`
}

// Shape ...
type Shape struct {
	ID      string   `xml:"id,attr" json:"id,omitempty"`
	Element string   `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Bounds  []Bounds `xml:"dc:Bounds" json:"bounds,omitempty"`
}

// TShape ...
type TShape struct {
	ID      string   `xml:"id,attr" json:"id,omitempty"`
	Element string   `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Bounds  []Bounds `xml:"Bounds" json:"bounds,omitempty"`
}
