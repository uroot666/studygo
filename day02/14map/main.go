package main

import "fmt"

// map

func main() {
	var m1 = make(map[string]int, 0) // 要估算好map容量，避免在程序运行期间动态扩容
	m1["理想"] = 9000
	m1["abc"] = 123
	fmt.Println(m1)

	fmt.Println(m1["理想"])
	v, ok := m1["abc"]
	if !ok {
		fmt.Println("无此key")
	} else {
		fmt.Println(v)
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	delete(m1, "abc")
	fmt.Println(m1)
}
