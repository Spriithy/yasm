package objects

import (
	"sync"
)

// Table is the central object type in this VM
// It is used to store anything, objects etc.
// It heavily relies on the system's map implementation
// to guaranty efficiency.
type Table struct {
	sync.RWMutex
	table map[string]Any
}

// CreateTable simply creates a new instance of a Table
// internal object with its Read/Write Mutex lock.
func CreateTable() *Table {
	return &Table{table: make(map[string]Any)}
}

// Class is used to determine the class name of any object
func (t *Table) Class() string {
	return "table"
}

// Copy returns a copy of the current table
func (t *Table) Copy() Any {
	return &Table{table: t.table}
}

// Zero returns the zero-value of the table type
func (t *Table) Zero() Any {
	return CreateTable()
}

// Read reads a pair value from the table using its key
// It either return nil, false if the key is not found
// Or value, true if it is
func (t *Table) Read(key string) (*Any, bool) {
	t.RLock()
	defer t.RUnlock()
	value, ok := t.table[key]
	return &value, ok
}

// Write is used to update a pair value in the map
// Or create it if the key is not mapped to anything
func (t *Table) Write(key string, pair Any) {
	t.Lock()
	defer t.Unlock()
	t.table[key] = pair
}

// Delete is used to delete an entry in a map
func (t *Table) Delete(key string) {
	t.Lock()
	defer t.Unlock()
	delete(t.table, key)
}

func (t *Table) String() string {
	if len(t.table) == 0 {
		return "{}"
	}
	str := ""
	t.RLock()
	for key, pair := range t.table {
		str += ", " + key + ": " + pair.String()
	}
	t.RUnlock()
	return "{" + str[2:] + "}"
}
