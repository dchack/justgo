package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
// 结构体
type Engine struct{}

// 结构体方法 对于传入的结构体参数有两种情况：1，传入指针；2，传入值
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL is %q\n", req.URL.Path)
	case "hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
