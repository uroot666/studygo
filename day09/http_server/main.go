package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/f1/", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
