package main

import "fmt"

// copy

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 3, 3)

	// 使用copy函数将切片中a1的元素复制到a3
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 将a1中索引为1的元素删掉
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)
}
