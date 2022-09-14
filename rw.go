package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	// synchronisation primitives
	var l sync.RWMutex
	var wg sync.WaitGroup
	readers := make(chan int)
	writers := make(chan int)

	// monitor access
	go func() {
		r, w := 0, 0
		for {
			select {
			case v := <-readers:
				r += v
			case v := <-writers:
				w += v
			}
			fmt.Printf("%s%s\n", strings.Repeat("R", r), strings.Repeat("W", w))
		}
	}()

	// launch reader and writer threads
	for i := 0; i < 100; i++ {
		// reader
		wg.Add(1)
		go func(id int, r chan int) {
			defer wg.Done()
			l.RLock()
			r <- 1
			r <- -1
			l.RUnlock()
		}(i, readers)

		// writer
		wg.Add(1)
		go func(id int, w chan int) {
			defer wg.Done()
			l.Lock()
			w <- 1
			w <- -1
			l.Unlock()
		}(i, writers)
	}
	wg.Wait()
}
