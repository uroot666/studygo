package main

import "fmt"

// 类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Printf("传进来的是字符串:%s\n", str)
	}
}

func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是字符串:", t)
	case int:
		fmt.Println("是int: ", t)
	case int64:
		fmt.Println("是int64: ", t)
	case bool:
		fmt.Println("是一个bool: ", t)
	}
}

func main() {
	assign2(false)
}
