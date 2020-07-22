package main

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "time"
)

type User struct {
	ID   int64
	Name string
	Age  int64 `gorm:"default:12"`
}

func main() {
	db, err := gorm.Open("mysql", "root:n246a369.@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{}) // 创建表

	// 创建记录
	u := User{
		Name: "kodt",
		Age:  20,
	}
	db.NewRecord(&u)      // 判断主键是否为空-->即表中是否有记录
	db.Debug().Create(&u) // 创建表中的一栏

}
