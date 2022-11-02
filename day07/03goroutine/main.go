package main

import (
	"fmt"
	"time"
)

// goroutine 对应的函数结束了，goroutine结束了
// main函数执行完了，由main函数创建的那些goroutine都结束了

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 10; i++ {
		go hello(i) //开启一个单独的goroutine去执行hello函数
	}
	fmt.Println("main")
	// main函数结束了，由main函数启动的goroutine也都结束了
	for i := 20; i < 30; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}
