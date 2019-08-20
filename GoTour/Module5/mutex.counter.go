package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (s *SafeCounter) Increment(key string) {
	s.mux.Lock()
	s.v[key]++
	s.mux.Unlock()
}

func (s *SafeCounter) Value(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Increment("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
