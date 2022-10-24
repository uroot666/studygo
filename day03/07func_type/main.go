package main

import "fmt"

func f1() {
	fmt.Println("Hello")
}

func f2() int {
	fmt.Println("f2")
	return 10
}

// 函数也可作为参数的类型
func f3(x func() int) {
	x()
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	f3(f2)

}
