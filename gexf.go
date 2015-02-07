package gexf4g

import (
	//"encoding/xml"
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
	DefualtEdgeType EdgeType      `xml:"defualtedgetype,attr,omitempty"`
	IdType          IdType        `xml:"id-type,attr,omitempty"`
	Mode            ModeType      `xml:"mode,attr,omitempty"`
	Attributes      AttributesDoc `xml:"attributes,omitempty"`
	Nodes           NodesDoc      `xml:"nodes"`
	Edges           EdgesDoc      `xml:"edges"`
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

// set meta
func (this *GexfDoc) SetMeta(_meta *MetaDoc) {
	this.Meta = _meta
}

//get meta
func (this *GexfDoc) GetMeta() *MetaDoc {
	//this.Meta = MetaDoc{}
	return this.Meta
}

//set meta lastmodefieddate
func (m *MetaDoc) SetLastModefiedDate(_date time.Time) {
	m.LastModifiedDate = _date
}

//get meta lastmodefieddate
func (m *MetaDoc) GetLastModefiedDate() time.Time {
	return m.LastModifiedDate
}

//set meta creator
func (m *MetaDoc) SetCreator(_creator string) {
	m.Creator = _creator
}

//get meta creator
func (m *MetaDoc) GetCreator() string {
	return m.Creator
}

//set meta keywords
func (m *MetaDoc) SetKeywords(_keywords string) {
	m.Keywords = _keywords
}

//get meta keywords
func (m *MetaDoc) GetKeyword() string {
	return m.Keywords
}

//set meta description
func (m *MetaDoc) SetDescription(_description string) {
	m.Description = _description
}

//get meta keywords
func (m *MetaDoc) GetDescription() string {
	return m.Description
}

//===========================================================
// Graph
//===========================================================
