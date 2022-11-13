package main

import (
	"flag"
	"fmt"
)

var name string
var age int
var married bool

func main() {
	name = *flag.String("name", "名字", "请输入名字")
	age = *flag.Int("age", 11, "请输入age")
	married = *flag.Bool("married", true, "结婚了吗")

	// flag.StringVar(&name, "name", "名字", "请输入名字")
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(name, age, married)
}
