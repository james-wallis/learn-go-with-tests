package sync

import "sync"

// Counter : A struct that represents a counter, holds a value that represents the current "Count"
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter : Constructor that initialises the Counter an returns its Pointer
func NewCounter() *Counter {
	return &Counter{}
}

// Inc : Increments the value by 1
func (c *Counter) Inc() {
	// Lock so no other goroutine can change the value, defer unlock
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value : Returns the current Counter value
func (c *Counter) Value() int {
	return c.value
}
