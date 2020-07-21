package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 中间件
// HandlerFunc
func index_han(c *gin.Context) {
	c.Get("name") // 取值
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

// 定义一个中间件m1
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	// 计时
	start := time.Now()
	c.Next() // 调用后续的处理函数
	//c.Abort() // 阻止调用后续的处理函数
	op := 1000000000
	for ; op > 0; op-- {

	}
	cost := time.Since(start)
	fmt.Printf("cose:%v\n", cost)
	fmt.Println("m1 out...")
	c.JSON(http.StatusOK, gin.H{
		"time": cost,
	})
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "kodt") // 存值
	//go funcxx(c.Copy())   只能使用拷贝的c
	c.Next()
	fmt.Println("m2 out...")

}

func authMiddlerware(doCheck bool) gin.HandlerFunc {
	// 连接数据库
	//其他准备
	return func(c *gin.Context) {
		if doCheck {
			// 存放具体的逻辑
			// if 登录用户
			// c.Next()
			//else
			// c.Abort()
		} else {
			//c.next()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(m1, m2, authMiddlerware(true)) // 全局注册中间件函数m1,m2
	r.GET("/index", m1, index_han)       // 部分注册中间件
	// 路由组注册中间件
	xxGroup := r.Group("/xx", authMiddlerware(false))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}
	r.Run(":9090")
}
