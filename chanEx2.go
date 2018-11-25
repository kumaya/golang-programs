package main

import (
	"fmt"
	"sync"
	"time"
)

func generate(m chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("put to chan: %d\n", i)
		m <- i
		time.Sleep(2 * time.Second)
	}
	close(m)
}

func consume(m <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case j, ok := <-m:
			if !ok {
				return
			}
			fmt.Printf("consume from chan: %d\n", j)
			time.Sleep(4 * time.Second)
		}
	}
}

func main() {
	msg := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go generate(msg, &wg)

	wg.Add(1)
	go consume(msg, &wg)
	wg.Wait()
}
