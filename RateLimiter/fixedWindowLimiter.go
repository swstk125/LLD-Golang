package main

import (
	"sync"
	"time"
)

type FixedWindow struct {
	count     int
	windowEnd time.Time
	mu        sync.Mutex
}

type FixedWindowLimiter struct {
	limit          int
	windows        map[string]*FixedWindow
	windowDuration time.Duration
	mu             sync.Mutex
}

func (f *FixedWindowLimiter) Allow(req string) bool {
	f.mu.Lock()
	window, ok := f.windows[req]
	if !ok {
		window = &FixedWindow{
			count:     0,
			windowEnd: time.Now().Add(f.windowDuration),
		}
		f.windows[req] = window
	}
	f.mu.Unlock()

	window.mu.Lock()
	defer window.mu.Unlock()

	now := time.Now()
	if now.After(window.windowEnd) {
		window.count = 0
		window.windowEnd = now.Add(f.windowDuration)
	}

	if window.count >= f.limit {
		return false
	}

	window.count++
	return true
}
