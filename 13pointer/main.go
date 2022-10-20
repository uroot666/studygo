package main

import "fmt"

// pointer

func main() {
	// 1. &:取地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)

	// 2 *: 根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	// make 和 new的区别
	// 1. make和new都是用来申请内存的
	// 2. new很少用，一般用来给基本数据类型申请内存，string/int ...，返回的是对应类型的指针
	// 3. make 是用来给slice map chan申请内存的，make 函数返回的是对于的类型本身

}
