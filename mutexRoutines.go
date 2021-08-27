package main

import (
	"fmt"
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

func foo(retry *int32) error {
	log.Println("inside foo1: ", *retry)
	atomic.AddInt32(retry, 1)
	//log.Println("inside foo2: ", *retry)
	time.Sleep(3 * time.Second)
	return fmt.Errorf("error from foo")
}

func bar(retry *int32) error {
	log.Println("inside bar1: ", *retry)
	atomic.AddInt32(retry, 1)
	//log.Println("inside bar2: ", *retry)
	time.Sleep(5 * time.Second)
	return fmt.Errorf("error from bar")
}

func main() {
	log.Println(runtime.NumGoroutine())
	var retry int32
	errChan := make(chan error)
	for retry < 4 {
		//log.Println("1: ", retry)
		go func() {
			errChan <- foo(&retry)
		}()
		go func() {
			errChan <- bar(&retry)
		}()
		log.Printf("+++ %v", <-errChan)
		go func() {
			drainChan(errChan)
		}()
		//log.Println("2: ", retry)
		time.Sleep(1 * time.Second)
	}
	time.Sleep(8 * time.Second)
	log.Println(runtime.NumGoroutine())
}

func drainChan(err <-chan error) {
	for e := range err {
		log.Printf("====== draining %v", e)
		return
	}
	//for {
	//	select {
	//	case e := <-err:
	//		log.Printf("====== draining %v", e)
	//		return
	//	}
	//}
}
