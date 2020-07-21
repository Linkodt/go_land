package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	kua := func(name string) (string, error) {
		return name + " nice guy", nil
	}
	t := template.New("hello.tmpl")
	// 告诉模板引擎，多了个自定义函数
	t.Funcs(template.FuncMap{
		"kua": kua,
	})
	// 解析模板
	_, err := t.ParseFiles("./hello.tmpl")

	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	name := "小王子"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Http server start failed,err=%v", err)
		return
	}
}
