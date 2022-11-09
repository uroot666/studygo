package main

import (
	"fmt"

	"github.com/uroot666/studygo/day09/split_string"
)

func main() {
	ret := split_string.Split("babcbef", "b")
	fmt.Printf("%#v\n", ret)
}
