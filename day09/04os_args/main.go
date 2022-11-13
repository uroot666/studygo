package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println(os.Args[0], os.Args[2])
	fmt.Printf("%T\n", os.Args)
}
