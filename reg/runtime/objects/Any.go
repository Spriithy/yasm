package objects

// Any is the base interface for every valid VM type object
type Any interface {
	Class() string
	String() string
	Copy() Any
}
