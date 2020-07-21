package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, _ := template.ParseFiles("./hello.tmpl")
	// 渲染模板
	err := t.Execute(w, "rush B")
	if err != nil {
		fmt.Println("errrrr:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Http server start failed,err=%v", err)
		return
	}

}
