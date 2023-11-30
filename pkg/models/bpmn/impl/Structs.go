package impl

type CoreID struct {
	ID string `xml:"id,attr,omitempty" json:"id"`
}

type CoreInnerID struct {
	ID string `xml:",innerxml,omitempty" json:"id"`
}

type BaseAttributes struct {
	ID   string `xml:"id,attr,omitempty" json:"id"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}
