package controller

import (
	"net/http"
	"../dao"
	"html/template"
)

// 用户登录
func Login(w http.ResponseWriter,r *http.Request){
	// 获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ :=dao.CheckUserNameAndPassword(username, password)

	if user.ID!=0{
		// 用户名和密码正确
		t:= template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	}else{
		t:= template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	}
}


func Regist(w http.ResponseWriter,r *http.Request){
	// 获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, _ :=dao.CheckUserName(username)

	if user.ID!=0{
		// 用户名和密码正确,用户已经存在
		t:= template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已经存在！")
	}else{
		dao.SavaUser(username, password, email)
		t:= template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

func CheckUserName(w http.ResponseWriter,r *http.Request){
	// 获取用户名
	username := r.PostFormValue("username")
	user, _ :=dao.CheckUserName(username)

	if user.ID!=0{
		// 用户名和密码正确
		w.Write([]byte("用户已经存在！"))
	}else{
		w.Write([]byte("<font style='color:green' >用户可用！</font>"))
	}
}