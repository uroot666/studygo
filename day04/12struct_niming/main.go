package main

import "fmt"

// 匿名字段
// 字段比较少爷比较简单的场景
// 不常用

type person struct {
	string
	int
}

func main() {
	p1 := person{
		"abc",
		111,
	}

	fmt.Println(p1)
	fmt.Println(p1.string)
}
