package gexf4g

import (
//"encoding/xml"
)

//class type
type ClassType uint8

const (
	NodeClassType ClassType = iota
	EdgeClassType
)

// attrtype
type AttrType uint8

const (
	IntegerAttrType AttrType = iota
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
	//Name xml.Name `xml:"attributes"`
	//Class     ClassType           `xml:"class, attr"`
	Class     string              `xml:"class,attr"`
	Attribute []*AttributeElement `xml:"attribute"`
}

//root > graph > attributes > attribute
type AttributeElement struct {
	//Name  xml.Name `xml:"attribute"`
	Id    int    `xml:"id,attr"`
	Title string `xml:"title,attr"`
	//Type    AttrType
	Type    string `xml:"type,attr"`
	Default string `xml:"deault,omitempty"`
	Options string `xml:"options,omitempty"`
}

//root > graph > nodes > node > attrvalues > attrvalues
//root > graph > edge > edge > attrvalues > attrvalues
type AttValuesDoc struct {
	XMLName struct{} `xml:"attvalues"`
	Values  []*AttrValueElement
}
type AttrValueElement struct {
	XMLName struct{} `xml:"attvalue"`
	//Name  xml.Name    `xml:"attvalue"`
	For   int         `xml:"for,attr"`
	Value interface{} `xml:"value,attr"`
}
