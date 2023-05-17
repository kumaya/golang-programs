package main

import (
	"fmt"
	"time"
)

type RateLimit interface {
	Allow() bool
	refill()
}

func main() {
	rl := NewTokenBucket(3, 2*time.Second)
	for i := 0; i < 20; i++ {
		if rl.Allow() {
			fmt.Println("allow")
		} else {
			fmt.Println("deny")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
