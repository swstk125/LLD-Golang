package main

import (
	"sync"
	"time"
)

type SlidingLog struct {
	timestamps []time.Time
	mu         sync.Mutex
}

type SlidingWindowLimiter struct {
	limit  int
	window time.Duration
	store  map[string]*SlidingLog
	mu     sync.Mutex
}

func NewSlidingWindowLimiter(limit int, window time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		limit:  limit,
		window: window,
		store:  make(map[string]*SlidingLog),
	}
}

func (l *SlidingWindowLimiter) Allow(key string) bool {
	l.mu.Lock()
	log, ok := l.store[key]
	if !ok {
		log = &SlidingLog{}
		l.store[key] = log
	}
	l.mu.Unlock()

	log.mu.Lock()
	defer log.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-l.window)

	// remove old timestamps
	idx := 0
	for idx < len(log.timestamps) && log.timestamps[idx].Before(cutoff) {
		idx++
	}
	log.timestamps = log.timestamps[idx:]

	if len(log.timestamps) >= l.limit {
		return false
	}

	log.timestamps = append(log.timestamps, now)
	return true
}
