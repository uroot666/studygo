package main

import (
	"fmt"
)

// 嵌套

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	address
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "abc",
		age:  111,
		address: address{
			province: "hubei",
			city:     "wuhan",
		},
	}
	fmt.Println(p1, p1.address.city)
	fmt.Println(p1.city)
}
