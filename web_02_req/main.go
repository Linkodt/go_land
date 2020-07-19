package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
	// http://localhost:8080/hello?name=1234
	fmt.Fprintln(w, "你发送的请求的地址是:",r.URL.Path) // /hello
	fmt.Fprintln(w, "你发送的请求地址后查询的字符串是:",r.URL.RawQuery)// name=1234
	fmt.Fprintln(w, "你发送的请求地址后请求头:",r.Header)
	fmt.Fprintln(w, "你发送的请求地址后请求头的Accept-Encoding:",r.Header["Accept-Encoding"])  // Accept-Encoding: [gzip, deflate, br]
	fmt.Fprintln(w, "你发送的请求地址后请求头的Accept-Encoding:",r.Header.Get("Accept-Encoding")) // Accept-Encoding: gzip, deflate, br
	// len := r.ContentLength
	// byte切片
	// body := make([]byte, len)
	// 将请求体内容读到body中
	// r.Body.Read(body)
	// 显示
	// fmt.Fprintln(w, "请求体中的内容有：",string(body))  // 请求体中的内容有： username=13&password=12
	/*
	r.ParseForm() // 这个方法调用后才可以调用 r.Form跟r.PostForm
	fmt.Fprintln(w,"请求参数有：", r.Form) // map[password:[123] username:[123]]   GET+POST
	fmt.Fprintln(w,"POST请求form表单参数有：", r.PostForm) // map[password:[123] username:[123]]   POST
	*/

	fmt.Fprintln(w, "URL中的user请求参数的值是：",r.FormValue("username"))
	fmt.Fprintln(w, "Form中的username请求参数的值是：",r.PostFormValue("username"))

	// map[
	// Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9] 
	// Accept-Encoding:[gzip, deflate, br] 
	// Accept-Language:[zh-CN,zh;q=0.9,en;q=0.8] 
	// Cache-Control:[max-age=0] 
	// Connection:[keep-alive] 
	// Sec-Fetch-Dest:[document] 
	// Sec-Fetch-Mode:[navigate] 
	// Sec-Fetch-Site:[none] 
	// Sec-Fetch-User:[?1] 
	// Upgrade-Insecure-Requests:[1] 
	// User-Agent:[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36]]


}

func main() {

	http.HandleFunc("/hello", handler)

	http.ListenAndServe(":8080", nil)

}