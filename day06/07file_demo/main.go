package main

import (
	"fmt"
	"os"
)

// 1. 文件对象的类型
// 2. 获取文件对象的详细信息

func main() {
	fileOjb, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("打开文件报错: %v", err)
		return
	}
	// 1. 文件对象的类型
	fmt.Printf("%T\n", fileOjb)
	// 2. 获取文件对象的详细信息
	fileInfo, err := fileOjb.Stat()
	if err != nil {
		fmt.Printf("获取文件信息失败: %v", err)
		return
	}
	fmt.Printf("文件大小: %dB\n", fileInfo.Size())
}
