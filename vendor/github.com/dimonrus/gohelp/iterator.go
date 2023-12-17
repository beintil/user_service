package gohelp

// Iterator struct
type Iterator struct {
	// Current cursor
	current int
	// Count of items in the iterator
	count int
}

// NewIterator Iterator constructor
func NewIterator(len int) *Iterator {
	return &Iterator{
		current: -1,
		count:   len,
	}
}

// Next Iterator next iteration
func (c *Iterator) Next() bool {
	if c.current >= c.count-1 {
		c.Reset()
		return false
	}
	c.current++
	return true
}

// Cursor Iterator Get cursor
func (c *Iterator) Cursor() int {
	return c.current
}

// Reset cursor
func (c *Iterator) Reset() *Iterator {
	c.current = -1
	return c
}

// Count Iterator Get count of items
func (c *Iterator) Count() int {
	return c.count
}

// SetCount Iterator Set count
func (c *Iterator) SetCount(count int) *Iterator {
	c.count = count
	return c
}
