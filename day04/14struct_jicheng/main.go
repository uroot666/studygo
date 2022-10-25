package main

import "fmt"

// 结构体模拟实现其它语言的继承

type animal struct {
	name string
}

// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

// 狗类
type dog struct {
	feet   uint8
	animal // animal 拥有的方法，dog此时也有了
}

// 给dog实现一个方法
func (d dog) wang() {
	fmt.Printf("%s ooooo....\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{name: "abc"},
		feet:   1,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
