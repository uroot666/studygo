package main

import "fmt"

func main() {
	s1 := []string{"北京", "深圳", "武汉"}

	// 调用append函数必须用原来的切片变量接收返回值
	// append 追加元素，原来的底层数组放不下的时候，go底层就会把底层数组换一个
	s1 = append(s1, "天津")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	ss := []string{"西安", "苏州"}
	s1 = append(s1, ss...) // ...表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
