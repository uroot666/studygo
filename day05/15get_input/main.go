package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时如果有空格

func userScan() {
	var s string
	fmt.Print("请输入: ")
	fmt.Scanln(&s)
	fmt.Printf("输入的内容: %v\n", s)
}

func userBufio() {
	var s string
	fmt.Print("请输入: ")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容: %v\n", s)
}

func main() {
	// userScan()
	userBufio()
}
