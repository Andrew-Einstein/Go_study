package day01

import (
	"fmt"
	"net/http"
)

// 简单web服务器
func main() {
	// 监听客户端请求
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:800", nil)

	// 响应客户端请求的回调函数
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>404!!!</h1>", r.URL.Path)

}
