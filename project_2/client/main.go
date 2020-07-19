package main

import (
	"fmt"
	"./login"
)

var Userid int
var Userpwd string

func main() {
	
	// 接收用户的选择
	var key int
	// 判断是否还继续显示菜单
	var loop = true

	for loop{
		fmt.Println("--------------欢迎登陆多人聊天系统---------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t  请选择(1-3):")

		fmt.Scanf("%d\n",&key)
		switch key{
		case 1 :
			fmt.Println("登录聊天室")
			loop = false // 跳出循环
		case 2 :
			fmt.Println("注册用户")
			loop = false
		
		case 3 :
			fmt.Println("退出系统")
			loop = false
			
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
	if key == 1 {
		// 用户要登录
		fmt.Println("请输入用户的id号:")
		fmt.Scanf("%d\n", &Userid)  // 要有斜杠
		fmt.Println("请输入用户的密码:")
		fmt.Scanf("%s\n", &Userpwd)
		// 登录函数写到另外一个文件， 比如login.go
		err := login.Login(Userid, Userpwd)
		if err != nil {
			fmt.Println("登录失败")
		}else{
			fmt.Println("登录成功")
		}

		
		}else{
			// fmt.Println("输入不是")
			
		}

}