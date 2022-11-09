package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:9090/hello/")
	if err != nil {
		fmt.Println("get url failed, err", err)
		return
	}

	// 从 resp中把服务端返回的数据读出来
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body failed, err:", err)
		return
	}
	fmt.Println(string(b))

}
