package main

import "fmt"

// 接口示例
// 不管什么牌子的车，都能跑

type car interface {
	run()
}

type falali struct {
	brand string
}

func (f falali) run() {
	fmt.Printf("%s 速度79\n", f.brand)
}

type baoshijie struct {
	brand string
}

func (b baoshijie) run() {
	fmt.Printf("%s 速度100\n", b.brand)
}

func drive(c car) {
	c.run()
}

func main() {
	var f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}

	drive(f1)
	drive(b1)
}
