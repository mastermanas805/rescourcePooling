package count

import "sync"

type Count struct {
	cnt int
	mu  sync.Mutex
}

func (c *Count) increment() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cnt++
	return c.cnt
}

func (c *Count) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cnt
}

var cnt = &Count{cnt: 0}

func GetCount() int {
	return cnt.getCount()
}

func IncrementCount() int {
	return cnt.increment()
}
