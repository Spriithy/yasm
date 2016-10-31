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

func ComparisonMode(t1, t2 BuiltinType) byte {
	const (
		HBYTE byte = 0x10
		HINTG byte = 0x20
		HREAL byte = 0x30

		TBYTE byte = 0x01
		TINTG byte = 0x02
		TREAL byte = 0x03
	)
	var MASK byte = 0x00
	if t1.IsNumeric() && t2.IsNumeric() {
		switch t1 {
		case BUILTIN_BYTE: MASK += HBYTE
		case BUILTIN_INTEGER: MASK += HINTG
		case BUILTIN_REAL: MASK += HREAL
		}

		switch t2 {
		case BUILTIN_BYTE: MASK += TBYTE
		case BUILTIN_INTEGER: MASK += TINTG
		case BUILTIN_REAL: MASK += TREAL
		}
		return MASK
	}

	return 0x00
}