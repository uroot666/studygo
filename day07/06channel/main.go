package main

import "fmt"

// var a []int
var b chan int // 需要指定通道中元素的类型

func main() {
	fmt.Println(b)     // nil
	b = make(chan int) // 不带缓冲区通道的初始化

	go func() {
		x := <-b
		fmt.Println("后台groutine从通道b中取到", x)
	}()
	b <- 10
	b = make(chan int, 16) //带缓冲区的通道的初始化
	fmt.Println(b)
}
