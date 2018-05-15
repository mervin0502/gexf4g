package gexf4g

//"encoding/xml"

// edge type
type EdgeType uint8
type DefaultEdgeType EdgeType

const (
	Directed EdgeType = iota
	Undirected
	Mutual
)

//root > graph > edges
type EdgesDoc struct {
	//Name  xml.Name `xml:"edge"`
	Count uint `xml:"count,attr,omitempty"`
	Edges []*EdgeElement
}

//root > graph > edges > edge
type EdgeElement struct {
	XMLName struct{} `xml:"edge"`
	//Name    xml.Name `xml:"edge"`
	//Type      EdgeType            `xml:"type, attr, omitempty"`
	Type   string  `xml:"type,attr,omitempty"`
	Label  string  `xml:"label,attr,omitempty"`
	Id     int     `xml:"id,attr"`
	Source int     `xml:"source,attr"`
	Target int     `xml:"target,attr"`
	Weight float32 `xml:"weight,attr,omitempty"`
	Test   int     `xml:"test,omitempty"`
	//TODO:why?
	//AttValues AttValuesDoc `xml:"attvalues,omitempty"`
	AttValues AttValuesDoc `xml:"attvalues,omitempty"`
}

func (e EdgeType) String() string {
	_t := "directed"
	switch e {
	case Directed:
		_t = "directed"
		break
	case Undirected:
		_t = "undirected"
	default:
		_t = "directed"
	}
	return _t
}
