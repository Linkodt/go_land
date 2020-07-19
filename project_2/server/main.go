package main

import (
	"net"
	"fmt"
	"../common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
)

func readPkg(conn net.Conn)(mes message.Message, err error){

		data := make([]byte, 8096)
		n, err := conn.Read(data[:4])
		if n != 4 || err != nil{
			err = errors.New("read pkg header error")
			return 
		}
		fmt.Println("读到的buf=",data[:4])
		// 根据读到的长度转成unit32类型 字节
		var pkgLen uint32
		pkgLen = binary.BigEndian.Uint32(data[0:4])
		// 上面的data都是存放着长度，下面的这个data此时接收的是数据包本身
		n, err = conn.Read(data[:pkgLen])
		if err != nil || n!=int(pkgLen) {
			err = errors.New("read pkg data error")
		return 
		}

		// 把data反序列化
		err = json.Unmarshal(data[:pkgLen], &mes) // &&&&&&这个符号记得加
		if err != nil {
			fmt.Println("json.Unmarsha err=", err)
			return 
		}

		return mes, err

}
// 处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message)(err error){
	// 先从mes中取出mes.data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=",err)
		return 
	}

	if loginMes.UserId == 100 && loginMes.UserPwd=="123456"{
		// 合法
	}else{
		// 不合法
	}

}


// 根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message)(err error){
	switch mes.Type{
	case message.LoginMesType:
		// 处理登录逻辑
		serverProcessLogin(conn,mes)
	case message.RegisterMesType:
		// 处理注册逻辑
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	// 读客户端发送的信息
	for{
		// 读取数据包封装成函数 readPkg
		mes,err := readPkg(conn)
		if err != nil {
			
			fmt.Println("readPkg err=",err)
			return 
		}
		fmt.Println(mes)
		serverProcessMes(conn, &mes)
	}
}


func main() {
	// 提示信息
	fmt.Println("服务器在8889端口监听。。。")
	listen,err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.listen err=",err)
		return 
	}
	// 一旦监听成功，等待客户端来链接服务器
	for{
		fmt.Println("等待客户端来链接服务器。。。")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err=",err)
		}

		// 一旦成功则启动一个协程和客户端保持通讯
		go process(conn)
	}
}