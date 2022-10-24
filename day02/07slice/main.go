package main

import "fmt"

// 切片slice

func main() {
	var s1 []int // 定义一个存放int类型元素的切片
	var s2 []string
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"北京", "上海"}
	fmt.Println(s1)
	fmt.Println(s2)

	// 长度和容量
	fmt.Printf("len(s1):%d, cap(s1):%d\n", len(s1), cap(s1))

	// 2. 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 基于一个数组切割，左包含右不包含
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)

	// 切片的容量是指底层数组的容量
	fmt.Printf("len(s3):%d, cap(s3):%d\n", len(s3), cap(s3))
	// 底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s4):%d, cap(s4):%d\n", len(s4), cap(s4))

	// 切片指向了一个底层数组
	// 切片的长度就是它元素的个数
	// 切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
}
