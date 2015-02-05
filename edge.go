package gexf4g

import (
	"encoding/xml"
)

// edge type
type EdgeType uint8
type DefaultEdgeType EdgeType

const (
	Directed EdgeType = itoa
	Undirected
	Mutual
)

//root > graph > edges
type EdgesDoc struct {
	Name  xml.Name `xml:"edge"`
	Count uint     `xml:"count, attr"`
	Edges []*EdgeElement
}

//root > graph > edges > edge
type EdgeElement struct {
	Name      xml.Name           `xml:"edge"`
	Type      EdgeType           `xml:"type, attr"`
	Label     string             `xml:"label, attr"`
	Id        int                `xml:"id, attr"`
	Source    int                `xml:"source, attr"`
	Target    int                `xml:"target, attr"`
	Weight    float32            `xml:"weight, attr"`
	AttValues []AttrValueElement `xml:"attvalues>attvalue, omitempty"`
}
