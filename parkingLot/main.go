package main

import (
	"fmt"
	"sync"
)

type Spot struct {
	available bool
	mu        sync.Mutex
}

func (s *Spot) bookSpot(wg *sync.WaitGroup) {
	s.mu.Lock()
	defer s.mu.Unlock()
	defer wg.Done()
	if s.available {
		fmt.Println("booking slot")
		s.available = false
		return
	}
	fmt.Println("slot unavailable")
}

func main() {
	S := &Spot{available: true}
	var wg sync.WaitGroup
	wg.Add(1)
	go S.bookSpot(&wg)
	wg.Add(1)
	go S.bookSpot(&wg)

	wg.Wait()

	fmt.Println(S.available)
}
