package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileOjb, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件报错: %v\n", err)
		return
	}
	log.SetOutput(fileOjb)

	for {
		log.Println("这是一条测试日志")
		time.Sleep(2 * time.Second)
	}
}
