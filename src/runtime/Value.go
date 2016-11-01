package runtime

import (
	"log"
	"fmt"
)

type Value struct {
	kind Type
	data interface{}
}

func Wrap(val interface{}) *Value {
	if val == nil {
		return &Value{kind: nullType, data: nil}
	}

	var k Type
	switch val.(type) {
	case bool: k = booleanType
	case uint8: k = byteType
	case int32: k = intType
	case int64: k = longType
	case float64: k = floatType
	case string: k = stringType
	default: k = nullType
	}

	return &Value{kind: k, data: val}
}

// Object Interface implementation

func (v Value) ToString() string {
	var d = v.data
	switch d.(type) {
	case uint8, int32, int64, float64:
		return fmt.Sprintf("%d", d)
	case bool:
		if d.(bool) {
			return "true"
		}
		return "false"
	case string:
		return fmt.Sprintf("%#v", d.(string))
	default:
		return "<?>"
	}
}

func (v Value) Equals(obj Object) bool {
	var val, ok = obj.(*Value)
	if !ok {
		return false
	}

	if val.data == v.data {
		return true
	}

	return false
}

func (v Value) Kind() ObjectKind {
	return ValueKind
}

func (v Value) Implements(k ObjectKind) bool {
	return k == v.Kind()
}

///////////

func (v *Value) Raw() interface{} {
	return v.data
}

func (v *Value) Byte() uint8 {
	var val, ok = v.data.(uint8)
	if !ok {
		log.Panicf("expected uint8 value, got %T", v.data)
	}
	return val
}

func (v *Value) Int() int32 {
	var val, ok = v.data.(int32)
	if !ok {
		log.Panicf("expected int32 value, got %T", v.data)
	}
	return val
}

func (v *Value) Long() int64 {
	var val, ok = v.data.(int64)
	if !ok {
		log.Panicf("expected int64 value, got %T", v.data)
	}
	return val
}

func (v *Value) Float() float64 {
	var val, ok = v.data.(float64)
	if !ok {
		log.Panicf("expected float64 value, got %T", v.data)
	}
	return val
}

func (v *Value) Bool() bool {
	var val, ok = v.data.(bool)
	if !ok {
		log.Panicf("expected bool value, got %T", v.data)
	}
	return val
}

func (v *Value) String() string {
	var val, ok = v.data.(string)
	if !ok {
		log.Panicf("expected string value, got %T", v.data)
	}
	return val
}

func (v *Value) IsNumeric() bool {
	switch (v.data).(type) {
	case uint8, int32, int64, float64:
		return true
	}
	return false
}

func (v *Value) IsNull() bool {
	return v.kind == nullType
}

func (v *Value) IsPointer() bool {
	return v.kind == pointerType
}
