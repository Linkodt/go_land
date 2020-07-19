package dao

import (
	"testing"
	"fmt"
)

func TestUser(t *testing.T){
	fmt.Println("测试userdao中的函数")
	t.Run("验证用户名或密码", testLogin)
	t.Run("验证用户名", testRegist)
	t.Run("保存用户", testSava)

}


func testLogin(t *testing.T){
	user,_:=CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是：",user)
}


func testRegist(t *testing.T){
	user,_:=CheckUserName("admin2")
	fmt.Println("获取用户信息是：",user)
}


func testSava(t *testing.T){
	SavaUser("admin2","123456","123@qq.com")
}