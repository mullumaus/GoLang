package main

import (
	"fmt"
	"sync"
	"time"
)

//SafeCounter is
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) increase(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

//increase the counter for the given key
func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.increase("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.value("somekey"))
}
