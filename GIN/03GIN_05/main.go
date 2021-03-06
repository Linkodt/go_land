package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Id    int
	Name  string
	Skill string
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, _ := template.ParseFiles("./hello.tmpl")
	// 渲染模板

	user := User{
		Id:    1,
		Name:  "kodt",
		Skill: "play game",
	}
	man := map[string]interface{}{
		"age":  18,
		"name": "kodt_1",
	}

	err := t.Execute(w, map[string]interface{}{ // 传递多个变量
		"u1": user,
		"m1": man,
	})
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
