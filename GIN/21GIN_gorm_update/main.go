package main

// 更新
import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "time"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:n246a369.@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 创建表
	db.AutoMigrate(&User{})

	var user User   // 声明变量
	db.First(&user) // select * from users where id=1
	fmt.Printf("user:%v\n", user)
	// 更新
	user.Name = "我是新的"
	user.Age = 222
	db.Save(&user) // 默认修改所有字段
	db.Model(&user).Update("name", "我是新的*2")

}
