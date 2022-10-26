package main

import "fmt"

// 引出接口的实例
// 接口是一种类型,是一种特殊变量，它规定了变量包含哪些方法
// 我不关心一个变量是什么类型，我只关心能调用它的什么方法

type cat struct{}

type dog struct{}

type person struct{}

type speaker interface {
	speak()
}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (p person) speak() {
	fmt.Println("啊啊啊~")
}

func da(x speaker) {
	// 接收一个参数，传进来什么，我就打什么
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

	var ss speaker
	ss = c1
	ss = d1
	ss = p1
	da(ss)

}
