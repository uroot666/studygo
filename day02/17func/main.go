package main

import "fmt"

// 函数
// 函数的定义
// 函数是一段代码的封装
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数有返回值的
func f3() string {
	ret := "f3"
	return ret
}

// 返回可以命名可以不命名
// 命名的返回值就相当于在函数中声明一个变量
// 使用命名返回值可以return后省略
func f4(x int, y int) (ret int) {
	ret = x + y
	return
}

// 多个返回值
func f5() (int, string) {
	return 1, "abc"
}

// 参数的类型简写
// 当参数中连续多个参数的类型一致时，可以将非最后一个参数的类型省略
func f6(x, y int) int {
	return x + y
}

// 可变长参数
// 可变长参数必须放在参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

// Go语言中函数没有默认参数这个概念

func main() {
	f7("b")
	f7("a", 1, 2, 4, 5)
}
