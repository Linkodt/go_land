package main

import (
	"net"
	"fmt"
)

func process(conn net.Conn){
	defer conn.Close()
	// 循环接收
	for{
		slice := make([]byte, 1024)
		n,err := conn.Read(slice)    //不发则堵塞，超时
		// fmt.Printf("服务器在等到客户端%v 发送信息\n",conn.RemoteAddr().String())
		if err != nil {
			fmt.Printf("Read出错了 err=%v",err)
		   	return 
		   }
		  // 显示内容
		fmt.Print(string(slice[:n]))   // 0-n
	}
}

func main() {
	
	fmt.Println("服务器开始监听")
	// 使用tcp协议    本地监听8888端口
	listen,err := net.Listen("tcp", "0.0.0.0:8888") //本地8888端口
	if err != nil {
		fmt.Printf("listen err! err=%v",err)
		return 
	}
	defer listen.Close() // 延时关闭

	// 循环等待
	for{
		// 等待客户端连接
		fmt.Println("等待客户端来链接")
		conn,err := listen.Accept()
		// defer conn.Close()
		if err != nil {
			fmt.Printf("Accept() err=", err)
		}else{
			fmt.Printf("Accept() suc con=%v 客户端地址=%v\n",conn,conn.RemoteAddr().String())
		}
		// 这里启协程为客户端服务
		go process(conn)
	}
	
	
	// fmt.Println("listen=",listen)
	// fmt.Printf("listen=%v", listen)
}