package runtime

type ObjectKind int

const (
	InvalidKind ObjectKind = iota

	ValueKind
	ListKind
	StructureKind
)

type Object interface {
	ToString() string
	Equals(Object) bool

	Kind() ObjectKind
	Implements(ObjectKind) bool
}

type Structure struct {
	name   string
	fields []Object
}