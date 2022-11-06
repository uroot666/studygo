package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock

var x = 0
var wg sync.WaitGroup

// var lock sync.Mutex
var rwlock sync.RWMutex

func read() {
	defer wg.Done()
	// lock.Lock()
	rwlock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwlock.Unlock()
}

func write() {
	defer wg.Done()
	// lock.Lock()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	// lock.Unlock()
	rwlock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
