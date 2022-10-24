package main

import "fmt"

func f1() {
	fmt.Println("Hello")
}

func f2() int {
	fmt.Println("f2")
	return 10
}

func ff(a, b int) int {
	return a + b
}

// 函数也可作为参数的类型
func f3(x func() int) {
	x()
}

// 函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	return ff
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	f3(f2)
	f7 := f5(f2)
	fmt.Printf("%T\n", f7)
	fmt.Println(f7(1, 2))
}
