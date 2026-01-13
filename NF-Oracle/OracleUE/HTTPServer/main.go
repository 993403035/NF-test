package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/.well-known/acme-challenge/", func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Path[len("/.well-known/acme-challenge/"):]
		fmt.Println("收到挑战请求，token:", token)

		// 模拟返回 token 作为响应体
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("challenge-response-for-token: " + "test-res"))
	})

	fmt.Println("监听端口 http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic("启动服务失败，可能需要 root 权限运行（80 端口）：" + err.Error())
	}
}
