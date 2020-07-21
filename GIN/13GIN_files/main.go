package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("./upload", func(c *gin.Context) {
		// 从请求中读取文件
		f, err := c.FormFile("f1") // 请求中获取携带的参数 一样的
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			// 将读取的文件保存到本地(服务端本地)
			dst := fmt.Sprintf("./%s", f.Filename)
			//dst := path.Join("./",f.Filename)
			err = c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
			if err != nil {
				return
			}
		}

	})

	r.Run(":9090")

}
