package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type User struct {
	Username string   `form:"username" json:"username"`
	Password string   `form:"password" json:"password"`
	Hobby    []string `form:"hobby" json:"hobby"`
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {

	c.Ctx.WriteString("用户中心")
}

func (c *UserController) AddUser() {
	c.TplName = "userAdd.html"
}

func (c *UserController) DoAddUser() {
	c.TplName = "user.html"
	user := c.GetString("username")
	pwd := c.GetString("password")
	hobby := c.GetStrings("hobby")
	fmt.Println(hobby)
	c.Ctx.WriteString("user" + user + "pwd:" + pwd)
}

func (c *UserController) EditUser() {
	c.TplName = "userEdit.html"
}
func (c *UserController) DoEditUser() {
	u := User{} // {Username:"123", Password:"453", Hobby:[]string{"2"}}
	if err := c.ParseForm(&u); err != nil {
		c.Ctx.WriteString("post提交失败")
		return
	}
	fmt.Printf("%#v", u)
	c.Ctx.WriteString("解析post数据成功")
}

func (c *UserController) GetUser() {
	u := User{
		Username: "张三",
		Password: "1234",
		Hobby:    []string{"1", "2"},
	}
	// 返回一个json
	c.Data["json"] = u
	c.ServeJSON()
}
