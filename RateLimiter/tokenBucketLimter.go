package main

import (
	"sync"
	"time"
)

/*
capacity
tokens
refillRate
lastRefillAt
*/

type TokenBucket struct {
	capacity     int
	tokens       int
	refillRate   int
	lastRefillAt time.Time
	mu           sync.Mutex
}

type TokenLimiter struct {
	buckets map[string]*TokenBucket
	mu      sync.Mutex
}

func NewTokenLimiter() *TokenLimiter {
	return &TokenLimiter{
		buckets: make(map[string]*TokenBucket),
	}
}

func (t *TokenLimiter) CreateBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		tokens:       capacity,
		refillRate:   refillRate,
		lastRefillAt: time.Now(),
	}
}

func (t *TokenLimiter) Allow(userId string) bool {
	t.mu.Lock()
	bucket, ok := t.buckets[userId]
	if !ok {
		bucket = t.CreateBucket(10, 2)
		t.buckets[userId] = bucket
	}
	t.mu.Unlock()

	// acquire lock
	bucket.mu.Lock()
	defer bucket.mu.Unlock()

	currTime := time.Now()
	elapsedTime := currTime.Sub(bucket.lastRefillAt).Seconds()
	updatedTokens := int(elapsedTime*float64(bucket.refillRate)) + bucket.tokens

	if updatedTokens > bucket.capacity {
		updatedTokens = bucket.capacity
	}

	bucket.tokens = updatedTokens
	bucket.lastRefillAt = currTime

	if bucket.tokens >= 1 {
		bucket.tokens--
		return true
	}
	return false
}
