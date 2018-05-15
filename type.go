package gexf4g

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

//String
func (i IdType) String() string {
	switch i {
	case Integer:
		return "interger"
	case String:
		return "string"
	default:
		return "string"
	}
	return "string"
}

//String
func (i ModeType) String() string {
	switch i {
	case Static:
		return "static"
	case Dynamic:
		return "dynamic"
	default:
		return "static"
	}
	return "static"
}
