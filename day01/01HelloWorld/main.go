package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	a := "abc"
	for _, r := range a {
		fmt.Printf("%c\n", r)
	}
}
