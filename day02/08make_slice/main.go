package main

import "fmt"

// make()函数创造切片
// 切片就是一个框，框住了一块连续的内存
// 切片属于引用类型，真正的数据都是保存在底层数组里的

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
