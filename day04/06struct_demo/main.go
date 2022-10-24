package main

import "fmt"

// 结构体占用一块连续的地址空间

type x struct {
	a int8 // 8bit => 1byte
	b int8
	c int8
	d string
}

func main() {
	m := x{
		int8(10),
		int8(11),
		int8(12),
		"abc",
	}

	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))
	fmt.Printf("%p\n", &(m.d))
}
