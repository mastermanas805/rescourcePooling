package api

import "sync"

type concurrentMap struct {
	mp map[string]State
	mu sync.Mutex
}

func (c *concurrentMap) set(key string, value State) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mp[key] = value
}

func (c *concurrentMap) get(key string) State {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.mp[key]
}

var mp = &concurrentMap{mp: make(map[string]State)}
