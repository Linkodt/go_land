package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("login.html", "index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		// 获取form表单的提交的数据
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		username := c.DefaultPostForm("username", "somebody")
		password := c.DefaultPostForm("password", "*******")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
		})
	})

	r.Run(":9090")
}
