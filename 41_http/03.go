package main

import(
	"fmt"
	"net/http"
)

// 创建处理器
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello World",r.URL.Path)
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// 创建路由
	http.ListenAndServe(":8080", mux)
}