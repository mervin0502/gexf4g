package gexf4g

import (
	"encoding/xml"
	"time"
)

//id type
var IdType uint8

const (
	Integer IdType = itoa
	String
)

//mode type
var ModeType uint8

const (
	Static ModeType = itoa
	Dynamic
)

//root element: gexf
type GexfDoc struct {
	Name    xml.Name `xml:"gexf"`
	Version float32  `xml:"version"`

	xmlNS          string `xml:"xmlns, attr"`
	xsi            string `xml:"xmlns:xsi, attr"`
	schemaLocation string `xml:"xsi:SchemaLocation, attr"`

	Meta  *MetaDoc  `xml:"meta, omitempty"`
	Graph *GraphDoc `xml:"graph"`
}

//gexf > meta
type MetaDoc struct {
	Name             xml.Name  `xml:meta`
	LastModifiedDate time.Time `xml:"lastmodifeddata, attr, omitempty"`
	Creator          string    `xml:"creator, omitempty"`
	Keywords         string    `xml:"keywords, omitempty"`
	Description      string    `xml:"description, omitempty"`
}

//root > Graph
type GraphDoc struct {
	Name            xml.Name       `xml:"graph"`
	DefualtEdgeType EdgeType       `xml:"defualtedgetype, omitempty,attr"`
	IdType          IdType         `xml:"id-type, omitempty,attr"`
	Mode            ModeType       `xml:"mode, omitempty,attr"`
	Attributes      *AttributesDoc `xml:attributes, omitempty`
	Nodes           NodesDoc       `xml:"nodes"`
	Edges           []*EdgeDoc     `xml:"edges>edge`
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
func (this *GexfDoc) SetMeta(_meta MetaDoc) {
	this.Meta = _meta
}

//get meta
func (this *GexfDoc) GetMeta() *MetaDoc {
	this.Meta = MetaDoc{}
	return this.Meta
}

//set meta lastmodefieddate
func (m *MetaDoc) SetLastModefiedDate(_date time.Time) {
	m.LastModifiedDate = _date
}

//get meta lastmodefieddate
func (m *MetaDoc) SetLastModefiedDate() time.Time {
	return m.LastModifiedDate
}

//set meta creator
func (m *MetaDoc) SetLastModefiedDate(_creator string) {
	m.Creator = _creator
}

//get meta creator
func (m *MetaDoc) SetLastModefiedDate() string {
	return m.Creator
}

//set meta keywords
func (m *MetaDoc) SetLastModefiedDate(_keywords string) {
	m.Keywords = _keywords
}

//get meta keywords
func (m *MetaDoc) SetLastModefiedDate() string {
	return m.Keywords
}

//set meta description
func (m *MetaDoc) SetLastModefiedDate(_description string) {
	m.Description = _description
}

//get meta keywords
func (m *MetaDoc) SetLastModefiedDate() string {
	return m.Description
}

//===========================================================
// Graph
//===========================================================
