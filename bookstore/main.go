package main

import (
	"net/http"
	"text/template"
	"./controller"
)

func IndexHandler(w http.ResponseWriter,r *http.Request){
	// 解析模板
	t := template.Must(template.ParseFiles("views/index.html"))

	// 执行
	t.Execute(w, "")
}



func main(){
	// 设置处理静态资源  将/static/ 替换成 views/static/
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))
	http.HandleFunc("/main", IndexHandler)

	// 去登陆
	http.HandleFunc("/login/",controller.Login)
	// 去注册
	http.HandleFunc("/regist/", controller.Regist)
	// checkusername
	http.HandleFunc("/checkUserName/", controller.CheckUserName)
	http.ListenAndServe(":8080",nil)
}