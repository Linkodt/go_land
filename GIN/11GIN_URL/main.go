package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/:age", func(c *gin.Context) {
		// 可以利用 name 跟 age参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	err := r.Run(":9090")
	if err != nil {
		return
	}
}
