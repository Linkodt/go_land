package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	msg := "这是index页面"
	t.Execute(w, msg)
}
func index2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err=%v\n", err)
		return
	}
	name := "小王子"
	t.Execute(w, name)
}
func home(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/index", index2)
	http.HandleFunc("/home", home)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server start failed,err:%v\n", err)
		return
	}
}
