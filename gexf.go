package gexf4g

import (
	"encoding/xml"
	"os"
	"time"
)

//id type
type IdType uint8

const (
	Integer IdType = iota
	String
)

//mode type
type ModeType uint8

const (
	Static ModeType = iota
	Dynamic
)

//root element: gexf
type GexfDoc struct {
	XMLName struct{} `xml:"gexf"`
	//Name    xml.Name `xml:"gexf"`
	Version float32 `xml:"version,attr"`

	XmlNS          string `xml:"xmlns,attr"`
	XSI            string `xml:"xmlns:xsi,attr"`
	SchemaLocation string `xml:"xsi:SchemaLocation,attr"`

	Meta  *MetaDoc  `xml:"meta,omitempty"`
	Graph *GraphDoc `xml:"graph"`
}

//gexf > meta
type MetaDoc struct {
	//Name             xml.Name  `xml:meta`
	LastModifiedDate time.Time `xml:"lastmodifeddata,attr,omitempty"`
	Creator          string    `xml:"creator,omitempty"`
	Keywords         string    `xml:"keywords,omitempty"`
	Description      string    `xml:"description,omitempty"`
}

//root > Graph
type GraphDoc struct {
	//Name            xml.Name       `xml:"graph"`
	DefualtEdgeType EdgeType `xml:"defualtedgetype,attr,omitempty"`
	IdType          IdType   `xml:"id-type,attr,omitempty"`
	Mode            ModeType `xml:"mode,attr,omitempty"`
	// NodeAttributes  *AttributesDoc `xml:"attributes,omitempty"`
	// EdgeAttributes  *AttributesDoc `xml:"attributes,omitempty"`
	Attributes []*AttributesDoc
	Nodes      *NodesDoc `xml:"nodes"`
	Edges      *EdgesDoc `xml:"edges"`
}

//initialize the gexf document
func NewGexfDoc() *GexfDoc {
	doc := &GexfDoc{
		Version: 1.2,

		XmlNS:          "http://www.gexf.net/1.2draft",
		XSI:            "http://www.w3.org/2001/XMLSchema-instance",
		SchemaLocation: "http://www.gexf.net/1.2draft http://www.gexf.net/1.2draft/gexf.xsd",
	}

	return doc
}

//===========================================================
// meta
//===========================================================
//SetMeta
func (gexf *GexfDoc) SetMeta(t time.Time, creator, keywords, description string) {
	gexf.Meta = &MetaDoc{
		LastModifiedDate: t,
		Creator:          creator,
		Keywords:         keywords,
		Description:      description,
	}
}

//===========================================================
// Graph
//===========================================================
//SetGraph
func (gexf *GexfDoc) SetGraph(edgeType EdgeType, idType IdType, modeType ModeType) *GraphDoc {
	gexf.Graph = &GraphDoc{
		DefualtEdgeType: edgeType,
		IdType:          idType,
		Mode:            modeType,
		Nodes:           &NodesDoc{Count: 0, Nodes: make([]*NodeElement, 0)},
		Edges:           &EdgesDoc{Count: 0, Edges: make([]*EdgeElement, 0)},
	}
	return gexf.Graph
}

//SetAttribute
func (graph *GraphDoc) SetAttribute(class ClassType) *AttributesDoc {

	if graph.Attributes == nil {
		graph.Attributes = make([]*AttributesDoc, 2)
	}
	var c string
	if class == NodeClassType {
		if graph.Attributes[0] == nil {
			c = "node"
			graph.Attributes[0] = &AttributesDoc{
				Class:     c,
				Attribute: make([]*AttributeElement, 0),
			}
		}
		return graph.Attributes[0]
	}
	if class == EdgeClassType {
		if graph.Attributes[1] == nil {
			c = "edge"
			graph.Attributes[1] = &AttributesDoc{
				Class:     c,
				Attribute: make([]*AttributeElement, 0),
			}
		}
		return graph.Attributes[1]
	}
	return nil
}

//GetAttribute
func (graph *GraphDoc) GetAttribute(class ClassType) *AttributesDoc {

	if class == NodeClassType {
		return graph.Attributes[0]
	} else if class == EdgeClassType {
		return graph.Attributes[1]
	} else {
		return nil
	}
}

//AddNode
func (graph *GraphDoc) AddNode(id int, label string) {
	graph.Nodes.Count = graph.Nodes.Count + 1
	e := &NodeElement{
		Id:    id,
		Label: label,
	}
	graph.Nodes.Nodes = append(graph.Nodes.Nodes, e)
}

//AddNodeWithAttribute
func (graph *GraphDoc) AddNodeWithAttribute(id int, label string, attrs map[int]interface{}) {
	graph.Nodes.Count = graph.Nodes.Count + 1
	attValues := AttValuesDoc{
		Values: make([]*AttrValueElement, len(attrs)),
	}
	for k, v := range attrs {
		att := &AttrValueElement{
			For:   k,
			Value: v,
		}
		attValues.Values = append(attValues.Values, att)
	}
	e := &NodeElement{
		Id:        id,
		Label:     label,
		AttValues: attValues,
	}
	graph.Nodes.Nodes = append(graph.Nodes.Nodes, e)
}

//AddEdge
func (graph *GraphDoc) AddEdge(id, src, tar int) *EdgeElement {
	graph.Edges.Count = graph.Edges.Count + 1
	e := &EdgeElement{
		Id:     id,
		Source: src,
		Target: tar,
	}
	graph.Edges.Edges = append(graph.Edges.Edges, e)
	return e
}

//AddEdge
func (graph *GraphDoc) AddEdgeWithAttribute(id, src, tar int, attrs map[int]interface{}) *EdgeElement {
	graph.Edges.Count = graph.Edges.Count + 1
	attValues := AttValuesDoc{
		Values: make([]*AttrValueElement, len(attrs)),
	}
	for k, v := range attrs {
		att := &AttrValueElement{
			For:   k,
			Value: v,
		}
		attValues.Values = append(attValues.Values, att)
	}
	e := &EdgeElement{
		Id:        id,
		Source:    src,
		Target:    tar,
		AttValues: attValues,
	}
	graph.Edges.Edges = append(graph.Edges.Edges, e)
	return e
}

//========================================================
//
//========================================================
//Marshal
func (gexf *GexfDoc) Marshal() ([]byte, error) {
	return xml.Marshal(gexf)
}

//MarshalToFile
func (gexf *GexfDoc) MarshalToFile(file string) {
	bytes, err := xml.Marshal(gexf)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	f.Write(bytes)
}
