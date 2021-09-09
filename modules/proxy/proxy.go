package proxy

// Counter ...
type Counter struct {
	Value int
}

// Inc ...
func (c *Counter) Inc() { c.Value++ }

// NewCounter ...
func NewCounter() *Counter { return &Counter{5} }
