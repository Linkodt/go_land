package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 方法二:结构体
type msg struct {
	Name    string `json:"name"`
	Age     int
	Message string
}

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//data := map[string]interface{}{
		//	"name":"kodt",
		//	"message":"hello world!",
		//	"age":18,
		//}
		data := gin.H{
			"name":    "kodt",
			"message": "hello world",
			"age":     18,
		}

		c.JSON(http.StatusOK, data)
	})
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			Name:    "小王子",
			Message: "hello world",
			Age:     20,
		}
		c.JSON(http.StatusOK, data) // json的序列化

	})
	err := r.Run(":9090")
	if err != nil {
		return
	}
}
