package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//r.GET("/hello",sayhello)
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"methoh": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	r.Run(":8080")
}

//GET 查询
//POST 创建
//PUT 更新
//DELETE 删除
