package main

import "fmt"

// defer
// defer 多用于函数结束之前释放资源（文件句柄、数据库链接

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("a") // defer 把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("b") // 一个函数可以有多个defer语句
	defer fmt.Println("c") // 多个defer语句按照先进后出的顺序执行
	fmt.Println("end")
}

func main() {
	deferDemo()
}
