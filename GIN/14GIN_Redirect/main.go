package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK,gin.H{
		//	"status":"ok",
		//})
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com") // 重定向
	})
	r.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "a",
		})
	})
	r.GET("/b", func(c *gin.Context) { //跳转到a
		c.Request.URL.Path = "/a" // 把请求的URL修改
		r.HandleContext(c)        // 继续后面的处理
	})
	r.Run(":9090")
}
