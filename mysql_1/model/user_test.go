package model

import (
	"testing"
	"fmt"
)

//TestMain 可以在测试函数之前做一些其他操作
func TestMain(m *testing.T){
	fmt.Println("测试开始：")
	// 通过m.Run()来执行测试
	// m.Run("测试开始：",TestUser)
	// 不用也可以
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	// 通过t.Run()来执行子测试函数
	// t.Run("测试添加用户：", testAddUser)
	// t.Run("测试查询用户：", testGetUserById)
	t.Run("测试查询用户：", testGetUsers)
	
}

func testAddUser(t *testing.T) {
	fmt.Println("测试添加用户：")
	user:=&User{}
	// 调用添加用户的方法
	user.AddUser()
	user.AddUser2()
}

func testGetUserById(t *testing.T){
	fmt.Println("测试查询一条记录")
	user:=User{
		ID:1,
	}
	// 调用获取User方法
	u,_ := user.GetUserById()
	// 返回一个结果指针
	fmt.Println("得到的User信息是：",*u)
}

// 获取所有user
func testGetUsers(t *testing.T){
	fmt.Println("测试查询所有记录：")
	user := &User{}
	// 调用获取所有User的方法
	us,_ := user.GetUsers()

	for i,v :=range us{
		fmt.Printf("第%d个用户是:%v", i+1,*v)
		fmt.Println("")
	}
}