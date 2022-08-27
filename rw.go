package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// readersCount >= 1, writersCount >= 1
var readersCount int = 25
var writersCount int = 10

var db sync.RWMutex
var wg sync.WaitGroup

// piece of data and its pointer
var value int
var p = &value

// sleep is for the output to be nicer, it is not a vital part of the algorithm
func sleep() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func reader(id int) {
	defer wg.Done()
	sleep()
	db.RLock()
	fmt.Printf("Reader%d : %d\n", id, *p)
	db.RUnlock()
}

func writer(id int) {
	defer wg.Done()
	sleep()
	db.Lock()
	*p = *p + 5
	fmt.Printf("\t\t\tWriter%d : %d\n", id, *p)
	db.Unlock()
}

func main() {
	for i := 0; i < writersCount; i++ {
		wg.Add(1)
		go writer(i + 1)
	}
	for i := 0; i < readersCount; i++ {
		wg.Add(1)
		go reader(i + 1)
	}
	wg.Wait()
}
