package controllers

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.Ctx.WriteString("新闻列表")
}

func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("增加新闻列表")
}

func (c *ArticleController) EditArticle() {
	id := c.GetString("id")
	c.Ctx.WriteString("修改新闻列表" + id)
}
