package main

import "fmt"

// 结构体

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个person类型的变量p
	var p person
	p.name = "test"
	p.age = 77
	p.gender = "男"
	p.hobby = []string{"1", "2", "3"}
	fmt.Println(p)

	// 访问变量p的字段
	fmt.Println(p.name)

	// 匿名结构体
	var s struct {
		name string
		age  int
	}

	s.name = "test2"
	s.age = 123
	fmt.Println(s)
}
