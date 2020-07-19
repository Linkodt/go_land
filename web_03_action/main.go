package main

import (
	"net/http"
	"html/template"
)

func testIf(w http.ResponseWriter,r *http.Request){
	// 解析模板文件
	t := template.Must(template.ParseFiles("index.html"))

	age := 17
	// 执行
	t.Execute(w, age>18)
}

func main() {
	http.HandleFunc("/testIf", testIf)

	http.ListenAndServe(":8080", nil)
}