package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting the application...")
	var waitGroup sync.WaitGroup
	// var data chan string
	data := make(chan string)

	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go worker(i, data, &waitGroup)
	}

	for i := 0; i < 15; i++ {
		data <- (strconv.Itoa(i))
	}
	close(data)
	waitGroup.Wait()
	fmt.Println("Stopping the application...")
}

func worker(workerId int, data chan string, wg *sync.WaitGroup) {
	fmt.Printf("Goroutine worker %d is starting...\n", workerId)
	time.Sleep(2 * time.Second)
	defer func(wg *sync.WaitGroup) {
		fmt.Printf("Destroying the worker: %d\n", workerId)
		wg.Done()
	}(wg)
	for {
		value, ok := <-data
		if !ok {
			fmt.Println("The channel is closed!")
			break
		}
		fmt.Printf("Worker: %d; channel value: %s\n", workerId, value)
		time.Sleep(time.Second * 1)
	}
}
