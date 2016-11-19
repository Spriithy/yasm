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

func (t *Table) Class() string {
	return "table"
}

func (t *Table) Copy() Any {
	return &Table{table: t.table}
}

func (t *Table) Read(key string) (*Any, bool) {
	t.RLock()
	defer t.RUnlock()
	value, ok := t.table[key]
	return &value, ok
}

func (t *Table) Write(key string, pair Any) {
	t.Lock()
	defer t.Unlock()
	t.table[key] = pair
}

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
