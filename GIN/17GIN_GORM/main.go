package main

// GORM操作mysql
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	u1 := &UserInfo{
		ID:     1,
		Name:   "kodt",
		Gender: "man",
		Hobby:  "play game",
	}
	//gorm.Open("mysql","user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True")
	db, err := gorm.Open("mysql", "root:n246a369.@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	u2 := UserInfo{ID: 2, Name: "kodt", Gender: "man", Hobby: "swim"}
	db.Create(&u1)
	db.Create(&u2) // 目前成功
	// 查询
	var u UserInfo
	db.First(&u) // u是第一个
	fmt.Printf("u:%v\n", u)
	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)
}
