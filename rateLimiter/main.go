package main

import (
	"log"
	"time"
)

type RateLimit interface {
	Allow() bool
}

func main() {
	log.Printf("===== TOKEN BUCKET =====")
	rl := NewTokenBucket(3, 2*time.Second)
	Run(rl)

	log.Printf("===== LEAKY BUCKET =====")
	rLB := NewLeakyBucket(3, 1*time.Second)
	Run(rLB)
}

func Run(rl RateLimit) {
	for i := 0; i < 10; i++ {
		if rl.Allow() {
			log.Println("allow")
		} else {
			log.Println("deny")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
