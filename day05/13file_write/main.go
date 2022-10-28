package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打开文件写内容

func write1() {
	fileOjb, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer fileOjb.Close()

	if err != nil {
		fmt.Printf("打开文件报错: %v\n", err)
		return
	}

	fileOjb.Write([]byte("test\n"))
	fileOjb.WriteString("test2\n")
}

func write2() {
	fileOjb, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer fileOjb.Close()

	if err != nil {
		fmt.Printf("打开文件报错: %v\n", err)
		return
	}
	wr := bufio.NewWriter(fileOjb)
	wr.Write([]byte("abc\n"))
	wr.WriteString("abc2\n")
	wr.Flush()
}

func main() {
	write2()
}
