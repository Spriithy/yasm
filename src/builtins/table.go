package builtins

import (
	"sync"

	uuid "github.com/Spriithy/go-uuid"
)

// The Table structure is used to map string keys to their
// Object pairs. Tables are used to implement objects.
type Table struct {
	id uuid.UUID
	sync.RWMutex
	table map[string]Object
}

// CreateTable creates a new instance of the Table object
func CreateTable() *Table {
	return &Table{id: uuid.NextUUID(), table: make(map[string]Object)}
}

// Set is used to update an Table's elements
func (t *Table) Set(key string, o Object) {
	t.Lock()
	t.table[key] = o.Copy()
	t.Unlock()
}

// Get is used to retrieve an Table's elements
// it returns NULL if the object is not present
func (t *Table) Get(key string) Object {
	t.RLock()
	defer t.RUnlock()
	if v, ok := t.table[key]; ok {
		return v
	}
	return NULL
}

// Size returns the number of elements in the table
func (t *Table) Size() Int {
	t.RLock()
	defer t.RUnlock()
	return Int(len(t.table))
}

// --- Object Interface ---

// Class returns a string describing the object's class
func (t *Table) Class() string {
	return "table"
}

// String returns a printable representation of the object
func (t *Table) String() string {
	if len(t.table) == 0 {
		return "{}"
	}

	str := ""
	t.RLock()
	for k, v := range t.table {
		str += ", " + k + ": "
		if v.Class() == "string" {
			str += v.(String).Format()
		} else {
			str += v.String()
		}
	}
	t.RUnlock()

	return "{" + str[2:] + "}"
}

// Equals test equality between two objects
func (t *Table) Equals(o Object) bool {
	if t.Class() != o.Class() {
		return false
	}
	return t.id == o.(*Table).id
}

// Copy returns an identical object yet somewhere else in memory
func (t *Table) Copy() Object {
	return &Table{id: t.id, table: t.table}
}
