package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	conn,err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("client dial err=", err)
		return  
	}
	// fmt.Println("conn=",conn," add=",conn.RemoteAddr().String())
	// 客户端发送单行数据，然后退出
	

		reader := bufio.NewReader(os.Stdin)  // 标准输入[终端]
		// 从终端读取一行用户输入，并准备发送给服务器
	for{
		line, err :=reader.ReadString('\n')
		strings.Trim(line," \r\n")
		if line == "exit"{return }  // 不知道为什么exit没退出
		if err != nil {
			fmt.Printf("readerString_err=%v",err)
	
		}
	
		// 发送给服务器
		_ ,err =conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err=",err)
		}
		// fmt.Printf("客户端发送了%d字节并退出",n)
	}



}