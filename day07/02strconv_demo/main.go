package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转换int64
	str := "1000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseint failed, err:", err)
		return
	}
	fmt.Printf("%T:%v\n", ret1, ret1)

	// 数字转换为字符串类型
	i := int32(83)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret2)
}
