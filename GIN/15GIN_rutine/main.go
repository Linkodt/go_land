package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由
func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.NoRoute(func(c *gin.Context) { // 所有未定义的地址
		c.JSON(http.StatusNotFound, gin.H{
			"method": "404",
		})
	})

	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
			})
		}
	})

	// 视频的首页和详情页
	// 路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/xx",
			})
		})
	}

	err := r.Run(":9090")
	if err != nil {
		return
	}
}
