package gexf4g

import (
	"encoding/xml"
)

//class type
var ClassType uint8

const (
	NodeClassType ClassType = itoa
	EdgeClassType
)

// attrtype
var AttrType uint8

const (
	IntegerAttrType AttrType = itoa
	LongAttrType
	DoubleAttrType
	FloatAttrTye
	BooleanAttrType
	ListAttrType
	StringAttrType
	AnyURIAttrType
)

//root > graph > attributes
type AttributesDoc struct {
	Name      xml.Name        `xml:"attributes"`
	Class     ClassType       `xml:"class, attr"`
	Attribute []*AttributeDoc `xml:"attribute"`
}

//root > graph > attributes > attribute
type AttributeElement struct {
	Name    xml.Name `xml:"attribute"`
	Id      int      `xml:"id, attr"`
	Title   string   `xml:"title, attr"`
	Type    AttrType `xml:"type, attr"`
	Default string   `xml:"deault, omitempty"`
	Options string   `xml:"options, omitempty"`
}

//root > graph > nodes > node > attrvalues > attrvalues
//root > graph > edge > edge > attrvalues > attrvalues
type AttrValueElement struct {
	Name  xml.Name `xml:"attvalue"`
	For   int      `xml:"for, attr"`
	Value int      `xml:"value, attr"`
}
