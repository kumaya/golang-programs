package main

import (
	"sync"
	"time"
)

// TokenBucket defines a rate limiter object.
type TokenBucket struct {
	capacity       int64         // defines the maximum number of tokens in the bucket
	token          int64         // defines the total number of tokens currently in the bucket
	refillRate     time.Duration // token refill rate; 1 per duration
	lastRefillTime time.Time
	lock           sync.Mutex
}

// NewTokenBucket creates a new rate limiter object.
func NewTokenBucket(cap int64, rate time.Duration) RateLimit {
	return &TokenBucket{
		capacity:       cap,
		token:          cap,
		refillRate:     rate,
		lastRefillTime: time.Now(),
		lock:           sync.Mutex{},
	}
}

// Allow checks if the rate limiter allows the request.
func (r *TokenBucket) Allow() bool {
	r.refill()
	if r.token > 0 {
		r.lock.Lock()
		r.token--
		r.lock.Unlock()
		return true
	}
	return false
}

// Refill implements token bucket rate limiting
func (r *TokenBucket) refill() {
	r.lock.Lock()
	defer r.lock.Unlock()
	timeElapsedSinceLastRefill := time.Since(r.lastRefillTime)
	tokensToFill := timeElapsedSinceLastRefill / r.refillRate
	partialTime := timeElapsedSinceLastRefill % r.refillRate
	if tokensToFill > 0 {
		r.token += tokensToFill.Nanoseconds()
		if r.token >= r.capacity {
			r.token = r.capacity
		}
		r.lastRefillTime = time.Now().Add(-partialTime)
	}
}
