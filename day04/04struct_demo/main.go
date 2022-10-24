package main

import "fmt"

// 结构体是值类型

type person struct {
	name, gender string
}

func f(x *person) {
	(*x).gender = "女"
}

func main() {
	var p person
	p.name = "test"
	p.gender = "男"
	f(&p)
	fmt.Println(p.gender)

	// 2. 结构体针1
	var p2 = new(person)
	(*p2).name = "test2"
	p2.gender = "女"

	// 结构体指针2
	var p3 = person{
		name:   "test3",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	// 使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序一致
	p4 := person{
		"abc",
		"男",
	}
	fmt.Printf("%#v\n", p4)
}
