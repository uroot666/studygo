package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readT(filename string) {
	fileObj, err := os.Open(filename)

	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
	}

	// 记得关闭文件
	defer fileObj.Close()

	// 读文件
	var tmp [2]byte
	for {

		n, err := fileObj.Read(tmp[:])
		if err != nil {
			fmt.Printf("读取错误, err: %v\n", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 2 {
			fmt.Println("读完了")
			return
		}
	}
}

// 利用bufio包读取文件
func readFromFilebyBufio(filename string) {
	fileObj, err := os.Open(filename)
	if err != nil {
		fmt.Printf("打开文件出错, %v\n", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()

	// 创建一个用来读取内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf(line)
			return
		}

		if err != nil {
			fmt.Printf("读取文件错误, err:%v", err)
		}
		fmt.Printf(line)
	}

}

func readFromFileByIoutil(filename string) {
	ret, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("读取错误, %v", err)
		return
	}
	fmt.Print(string(ret))

}

func main() {
	readT("t.txt")
	fmt.Println("------------")
	readFromFilebyBufio("t.txt")
	fmt.Println("------------")
	readFromFileByIoutil("t.txt")
}
