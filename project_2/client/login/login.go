package login
// 登录
import (
	"fmt"
	"net"
	"../../common/message"
	"encoding/json"
	"encoding/binary"
)

func Login(Userid int, Userpwd string)(err error){ // 完成登录
	// 定协议
	// fmt.Printf("userid=%d  userpwd=%s\n", Userid,Userpwd)
	// return nil
	
	// 1. 链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=",err)
		return 
	}
	defer conn.Close()

	// 2. 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	// 3. 创建一个LoginMes结构体
	var loginMes message.LoginMes // message.data
	loginMes.UserId = Userid
	loginMes.UserPwd = Userpwd

	// 4. 讲loginmes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err=", err)
		return 
	}
	// 5. 把data赋值给mes.Data字段
	mes.Data = string(data)
	// 6. 将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return 
	}

	// data此时是[]byte切片
	// 7.1 先发data长度
	// 先获取data长度，然后转成一个表示长度的切片
	var pkgLen uint32  // 2^32 - 1
	pkgLen = uint32(len(data))
	data_Len := make([]byte, 4)
	binary.BigEndian.PutUint32(data_Len[0:4], pkgLen) // 将长度转成bytes了

	// 发送长度
	n, err := conn.Write(data_Len)
	if n!=4 || err != nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return 
	}
	// fmt.Println("客户端发送消息的长度成功")
	// fmt.Println("客户端发送的长度为",len(data))

	// 发送消息本身
	_, err = conn.Write(data)
	if n!=4 || err != nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return 
	}
	
	return 
}