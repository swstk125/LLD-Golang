package main

import (
	"sync"
	"time"
)

type leakyBucket struct {
	capacity   int
	reqCount   int
	leakRate   int
	lastLeakAt time.Time
	mu         sync.Mutex
}

type LeakyLimiter struct {
	buckets map[string]*leakyBucket
	mu      sync.Mutex
}

func NewLeakyLimiter() *LeakyLimiter {
	return &LeakyLimiter{
		buckets: make(map[string]*leakyBucket),
	}
}

func (l *LeakyLimiter) CreateBucket(cap int, rate int) *leakyBucket {
	return &leakyBucket{
		capacity:   cap,
		reqCount:   0,
		leakRate:   rate,
		lastLeakAt: time.Now(),
	}
}

func (l *LeakyLimiter) Allow(reqId string) bool {
	l.mu.Lock()
	bucket, ok := l.buckets[reqId]

	if !ok {
		bucket = l.CreateBucket(5, 2)
		l.buckets[reqId] = bucket
	}
	l.mu.Unlock()

	bucket.mu.Lock()
	defer bucket.mu.Unlock()

	currTime := time.Now()
	elapsedTime := currTime.Sub(bucket.lastLeakAt).Seconds()

	reqCompleted := bucket.reqCount - int(elapsedTime*float64(bucket.leakRate))

	if reqCompleted < 0 {
		bucket.reqCount = 0
	} else {
		bucket.reqCount = reqCompleted
	}
	bucket.lastLeakAt = currTime

	if bucket.reqCount < bucket.capacity {
		bucket.reqCount++
		return true
	}
	return false
}
