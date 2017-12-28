package gexf4g

//"encoding/xml"

//root > graph > nodes
type NodesDoc struct {
	//Name  xml.Name `xml:nodes`
	Count uint `xml:"count,attr,omitempty"`
	Nodes []*NodeElement
}

//root > graph > nodes>node
type NodeElement struct {
	XMLName struct{} `xml:"node"`
	//Name xml.Name `xml:node`
	//Id        IdType              `xml:"id, attr"`
	Id        int          `xml:"id,attr"`
	Label     string       `xml:"label,attr"`
	AttValues AttValuesDoc `xml:"attvalues,omitempty"`
}
