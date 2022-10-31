package main

import (
	"fmt"
	"os"
)

func f2() {
	fileOjb, err := os.OpenFile("./xx.txt", os.O_RDWR|os.O_CREATE, 0644)
	defer fileOjb.Close()
	if err != nil {
		fmt.Printf("打开文件错误: %v\n", err)
		return
	}
	fileOjb.Seek(4, 0)

	var s []byte
	s = []byte{'c'}
	fileOjb.Write(s)

	var ret [1]byte
	n, err := fileOjb.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed, err: %v", err)
	}
	fmt.Println(string(ret[:n]))
}

func main() {
	f2()
}
