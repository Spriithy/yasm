package builtins

type BuiltinType string

const (
	BUILTIN_NULL = "Null"
	BUILTIN_BOOLEAN = "Boolean"
	BUILTIN_BYTE = "Byte"
	BUILTIN_INTEGER = "Integer"
	BUILTIN_REAL = "Real"
	BUILTIN_STRING = "String"
)

func TypeCheck(obj, other Object) Boolean {
	return obj.TypeOf() == other.TypeOf()
}

func (t BuiltinType) IsNative() Boolean {
	return t == BUILTIN_NULL || t == BUILTIN_STRING || t == BUILTIN_BOOLEAN ||
		t == BUILTIN_BYTE || t == BUILTIN_INTEGER || t == BUILTIN_REAL
}

func (t BuiltinType) IsNumeric() Boolean {
	return t == BUILTIN_BYTE || t == BUILTIN_INTEGER || t == BUILTIN_REAL
}