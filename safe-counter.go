package main

import (
	"fmt"
	"time"
	"sync"
)

type SafeCounter struct {
	mu sync.Mutex
	value map[string] int
}

func (s *SafeCounter) Inc(key string) {
	s.mu.Lock()
	s.value[key]++
	s.mu.Unlock()
}

func (s *SafeCounter) Value(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.value[key]
}

func main() {
	c := SafeCounter{value: make(map[string] int)}
	for i := 0; i < 1000; i++ {
		c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}