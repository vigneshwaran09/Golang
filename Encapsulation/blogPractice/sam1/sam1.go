package sam1

// SimpleCounter - a simple counter.
type SimpleCounter struct {
	counter int
}

// NewCounter - creates a new counter instance.
func NewCounter() *SimpleCounter {
	return &SimpleCounter{
		counter: 0,
	}
}

// N - return current counter's value.
func (c *SimpleCounter) N() int { return c.counter }

// Increment - increases the counter by 1 without considering the risk of race condition.
func (c *SimpleCounter) Increment() { c.counter++ }

// Reset - resets the counter.
func (c *SimpleCounter) Reset() { c.counter = 0 }
