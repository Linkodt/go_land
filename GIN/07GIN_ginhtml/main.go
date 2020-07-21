package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 静态文件
// html上用到的样式文件.css js文件 图片
func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/xxx", "./static")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("templates/index.tmpl") // 模板解析
	r.GET("/index", func(c *gin.Context) {
		// http请求
		c.HTML(http.StatusOK, "index.tmpl", gin.H{ // 模板渲染
			"title": "<a href='https://www.baidu.com'>kodt.com</a>",
		})
	})
	err := r.Run(":9090")
	if err != nil {
		return
	}

}
