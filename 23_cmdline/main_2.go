package main

// 用于接收命令行参数
import (
	"fmt"
	"flag"
)

func main() {
	// 定义变量，接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	// &user 接收用户命令行中输入的 -u后面的参数值
	// "u" 就是-u指定参数
	// "", 默认值
	// "用户名，默认为空" 说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	// 这个是转换操作 必须有
	flag.Parse()

	fmt.Printf("user=%v\npwd=%v\nhost=%v\nport=%v", user,pwd,host,port)

}