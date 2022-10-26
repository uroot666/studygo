package main

import "fmt"

// 使用值接受者和指针接收者的区别
// 使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存
// 指针接收者实现接口只能存结构体指针的变量

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// // 使用值接收者
// func (c cat) move() {
// 	fmt.Println("走猫步...")
// }

// func (c cat) eat(food string) {
// 	fmt.Printf("猫吃%s...\n", food)
// }

// 用值接收者
func (c *cat) move() {
	fmt.Println("走猫步...")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {

	var a1 animal
	c1 := cat{"tome", 4}

	c2 := &cat{"jia", 6}

	a1 = &c1
	a1 = c2
	a1.eat("123")
}
