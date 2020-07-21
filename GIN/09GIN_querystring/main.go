package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		name := c.Query("name") // 通过query获取请求中携带的name参数
		age := c.Query("age")
		//name := c.DefaultQuery("query","somebody") // 取不到就用指定的默认值
		//name,ok := c.GetQuery("query") // 娶不到则返回false
		//if !ok {
		//	name="somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})

	})
	r.Run(":9090")
}
