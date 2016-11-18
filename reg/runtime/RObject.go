package runtime

// RObject is the internal representation of what can be an Object
type RObject interface {
	ValueOf() interface{}
	TypeOf() RType
	String() RString
	SizeOf() RInt
}
