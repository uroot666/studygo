package main

import (
	"encoding/json"
	"fmt"
)

// 结构体和json

// 1. 序列化 把go语言中的结构体变量 -> json格式的字符串
// 2. 反序列化 json格式的字符串 -> go语言能识别的结构体变量

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "abc",
		Age:  111,
	}

	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err: %v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))

	// 反序列化
	str := `{"name":"liangbiao", "age":11}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%T, %#v\n", p2, p2)
}
