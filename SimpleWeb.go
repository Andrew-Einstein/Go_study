package main

import (
	"fmt"
	"net/http"
)

// 简单web服务器
func main() {
	// 监听客户端请求
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)

	// 响应客户端请求的回调函数
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello kiki", r.URL.Path)

}
