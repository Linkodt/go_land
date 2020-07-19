package main

import (
	_"fmt"
	"net/http"
	"html/template"
)

func testTemplate(w http.ResponseWriter,r *http.Request){
	// 解析模板文件
	// t,_ := template.ParseFiles("index.html")

	// 通过Must函数让go帮我们自动处理异常
	t:=template.Must(template.ParseFiles("index.html","index2.html"))

	// t.Execute(w, "hello template")
	t.ExecuteTemplate(w, "index2.html", "we go to index2.html")
}


func main() {
	http.HandleFunc("/hello", testTemplate)

	http.ListenAndServe(":8080", nil)
}