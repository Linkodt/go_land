package model

import (
	"../utils"
	"fmt"
)

// User 结构体
type User struct{
	ID		 	int
	Username 	string
	Password 	string
	Email 		string
}

//AddUser 添加User的方法一
func (user *User) AddUser() error{
	// 1. sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	// 2. 预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err!=nil{
		fmt.Println("预编译出现异常:",err)
		return err
	}

	// 3. 执行
	_,err2 := inStmt.Exec("admin","123456","admin@qq.com")
	if err2 != nil {
		fmt.Println("执行出现异常",err2)
		return err2
	}
	return nil
}

//AddUser2 添加User的方法二
func (user *User) AddUser2() error{
	// 1. sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	// 2. 执行
	_,err := utils.Db.Exec(sqlStr,"admin2","654321","admin2@qq.com")
	if err != nil {
		fmt.Println("执行出现异常",err)
		return err
	}
	return nil
}

//GetUserById 根据用户id从数据库中查询一条记录
func (user *User) GetUserById()(*User,error){
	// SQL语句
	sqlStr := "select id,username,password,email from users where id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, user.ID)
	// 声明
	var id 		 int
	var username string
	var password string
	var email 	 string
	err := row.Scan(&id, &username,&password,&email)
	if err != nil{
		return nil,err
	}

	u:=&User{
		ID:			id,
		Username:	username,
		Password:	password,
		Email:		email,
	}
	return u,nil
}

// GetUsers 获取数据库中所有的记录
func (user *User) GetUsers() ([]*User,error){
	sqlStr := "select id,username,password,email from users"

	rows,err := utils.Db.Query(sqlStr)

	if err != nil {

		return nil,err 
	}

	var Users []*User

	for rows.Next(){
		var id 		 int
		var username string
		var password string
		var email 	 string
		err := rows.Scan(&id, &username,&password,&email)
		if err != nil{
			return nil,err
		}

		u:=&User{
			ID:			id,
			Username:	username,
			Password:	password,
			Email:		email,
		}
		Users = append(Users, u)
	}
	return Users,nil
}