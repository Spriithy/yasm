package runtime

// RType is used to identify both builtin and user defined data types
type RType uint

const (
	// TNull is the Null type constant
	TNull = RType(iota)

	// TBool is the bool type constant
	TBool

	// TByte is the byte type constant
	TByte

	// TInt is the int type constant
	TInt

	// TUInt is the uint type constant
	TUInt

	// TFloat is the float type constant
	TFloat

	// TString is the string type constant
	TString

	// TTuple is the tuple type constant
	TTuple

	// TFunction is the function type constant
	TFunction

	// TThread is the thread type constant
	TThread

	// Anything above userData is a user defined type
	userData
)

// TypeName is used to map a type and its name only for builtin types
var TypeName = map[RType]RString{
	TNull:     StringOf("null"),
	TBool:     StringOf("bool"),
	TByte:     StringOf("byte"),
	TInt:      StringOf("int"),
	TUInt:     StringOf("uint"),
	TFloat:    StringOf("float"),
	TString:   StringOf("string"),
	TTuple:    StringOf("tuple"),
	TFunction: StringOf("function"),
	TThread:   StringOf("Thread"),
}
