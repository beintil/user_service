package gosql

// TableModifier func to modify table
type TableModifier func(t *Table)

// TableModeler list of TableModifier
type TableModeler []TableModifier

// Add define next table population method
func (t TableModeler) Add(m ...TableModifier) TableModeler {
	return append(t, m...)
}

// Prepare table
func (t TableModeler) Prepare(tb *Table) {
	for _, modifier := range t {
		modifier(tb)
	}
	return
}
