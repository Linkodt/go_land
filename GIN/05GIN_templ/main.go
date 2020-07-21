package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("t.tmpl", "ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed err:%v\n", err)
		return
	}
	// 渲染模板
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/tempDemo", demo1)
	http.ListenAndServe(":9090", nil)
}
