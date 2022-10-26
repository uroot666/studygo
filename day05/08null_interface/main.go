package main

import "fmt"

// 空接口

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "周林"
	m1["age"] = 9000
	m1["merried"] = true
	m1["hobby"] = [...]string{"1", "2", "3"}
	fmt.Println(m1)
	fmt.Printf("%T, %v\n", m1["name"], m1["name"])
}
