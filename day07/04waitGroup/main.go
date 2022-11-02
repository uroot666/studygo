package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano()) //随机数种子，保证每次不一样
		r1 := rand.Int()
		r2 := rand.Intn(10) // 0<= x < 10
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	fmt.Println("main")
	wg.Wait() // 等待wg的计数器减为0
}
