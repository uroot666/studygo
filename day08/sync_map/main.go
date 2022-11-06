package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int, 10)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

var m2 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)         // 必须使用sync.map内置的Store方法设置键值对
			value, _ := m2.Load(key) // 必须使用sync.map内置的Load方法获取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
