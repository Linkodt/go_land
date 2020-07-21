package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"username" `
	Password string `form:"password" json:"password" `
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := UserInfo{
		//	username,
		//	password,
		//}
		var u UserInfo //声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Println("%v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})
	r.POST("/form", func(c *gin.Context) {
		var u UserInfo //声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Println("%v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})
	err := r.Run(":9090")
	if err != nil {
		return
	}
}
