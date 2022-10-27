package main

import (
	"fmt"

	"github.com/uroot666/studygo/day05/11import_demo/calc"
)

// 自动执行，不能传参也没有返回值
func init() {
	fmt.Println("初始化...")
}

func main() {
	fmt.Println(calc.Add(1, 1))
}
