package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w, "测试Http协议！")
}

func main() {
	// 处理器处理请求
	http.HandleFunc("/http", Handler)

	// 路由
	http.ListenAndServe(":8080", nil)
	
}