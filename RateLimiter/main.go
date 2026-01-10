package main

type RateLimiter interface {
	Allow(string) bool
}
