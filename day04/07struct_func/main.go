package main

import "fmt"

// 构造函数

type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

// 构造函数
// 返回的是结构体还是结构体指针
// 当结构体比较大的的时候尽量使用结构体指针，减少程序的内存开销
func newPersion(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p1 := newPersion("a", 11)
	p2 := newPersion("b", 12)
	fmt.Println(p1, p2)
}
