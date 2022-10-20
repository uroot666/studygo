package main

import "fmt"

// make 和 slice 组合
// 切片和map一定要make初始化

func main() {
	// 元素类型为map的切片
	var s1 = make([]map[int]string, 1, 10)

	s1[0] = make(map[int]string, 1)
	s1[0][100] = "a"
	fmt.Println(s1)

	// 值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["b"] = []int{10, 20}
	fmt.Println(m1)
}
