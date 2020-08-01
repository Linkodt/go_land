package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var (
	DB *gorm.DB
)

func initMySQL() (err error) {
	DB, err = gorm.Open("mysql", "root:n246a369.@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
		return err
	}
	err = DB.DB().Ping()
	return nil
}

func main() {
	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	// 创建表 连接起来
	DB.AutoMigrate(&Todo{})
	r := gin.Default()
	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 增删改查
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端发送请求
			// 1. 从请求中拿出数据
			var todo Todo
			c.BindJSON(&todo) // ?
			// 2. 存入数据库
			// 3. 返回响应
			if err := DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				//panic(err)
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 查看所有待办
		v1Group.GET("/todo", func(c *gin.Context) {
			// 查询todo这个表里的所有数据
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"eooro": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		// 查看一个待办
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改某一个待办
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			// 获取值
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 删除某一个
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}
	err = r.Run(":9090")
	if err != nil {
		panic(err)
	}

}
