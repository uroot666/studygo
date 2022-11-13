package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// var notify bool
var exitChan chan bool = make(chan bool, 1)

func f() {
	defer wg.Done()
	for {
		fmt.Println("in..")
		time.Sleep(time.Millisecond * 500)
		// if notify {
		// 	break
		// }

		select {
		case <-exitChan:
			return
		default:
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// notify = true
	// wg.Done()
	exitChan <- true
	wg.Wait()
}
