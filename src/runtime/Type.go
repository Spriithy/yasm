package runtime

type Type int

const (
	nullType Type = iota
	pointerType

	booleanType

	byteType
	intType
	longType

	floatType

	stringType


	listType
)