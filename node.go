package gexf4g

import (
	"encoding/xml"
)

//root > graph > nodes
type NodesDoc struct {
	Name  xml.Name `xml:nodes`
	Count uint     `xml:"count, attr"`
	Nodes []*NodeElement
}

//root > graph > nodes>node
type NodeElement struct {
	Name      xml.Name      `xml:node`
	Id        IdType        `xml:"id, attr"`
	Label     string        `xml:"label, attr"`
	AttValues []AttValueDoc `xml:"attvalues>attvalue, omitempty"`
}
