package main

import (
	"sync"
	"time"
)

// TODO:
// add implementation of leaky bucket rate limiter

type LeakyBucket struct {
	capacity        int64         // defines the maximum number of tokens in the bucket
	token           int64         // defines the total number of tokens currently in the bucket
	emptyRate       time.Duration // rate at which the requests are removed from bucket.
	lastProcessTime time.Time
	lock            sync.Mutex
}

func NewLeakyBucket(cap int64, rate time.Duration) RateLimit {
	return &LeakyBucket{
		capacity:        cap,
		token:           0,
		emptyRate:       rate,
		lastProcessTime: time.Now(),
		lock:            sync.Mutex{},
	}
}

func (r *LeakyBucket) Allow() bool {
	r.empty()
	if r.token < r.capacity {
		r.lock.Lock()
		r.token++
		r.lock.Unlock()
		return true
	}
	return false
}

func (r *LeakyBucket) empty() {
	r.lock.Lock()
	defer r.lock.Unlock()
	timeElapsedSinceLastProcessed := time.Since(r.lastProcessTime)
	tokensToRemove := timeElapsedSinceLastProcessed / r.emptyRate
	delta := timeElapsedSinceLastProcessed % r.emptyRate
	r.token -= tokensToRemove.Nanoseconds()
	if r.token < 0 {
		r.token = 0
	}
	r.lastProcessTime = time.Now().Add(-delta)
}
