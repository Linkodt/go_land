package controllers

import (
	"github.com/astaxie/beego"
)

type GoodController struct {
	beego.Controller
}

func (c *GoodController) Get() {
	c.Data["Title"] = "hello world"
	c.TplName = "goods.tpl"
}
