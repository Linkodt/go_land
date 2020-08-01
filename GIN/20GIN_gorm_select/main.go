package main

// 查询
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
	// 创建
	//u1 := User{Name:"kodt_1",Age: 15}
	//u2 := User{Name:"kodt_1",Age: 18}
	//db.Create(&u2)
	//db.Create(&u1)

	// 4.查询
	var user User   // 声明变量
	db.First(&user) // select * from users where id=1
	fmt.Printf("user:%v\n", user)
	// db.Take() 随机取一条
	// db.Last() 根据主键查询最后一条
	// ad.Find() 查询所有
	// db.First(&user,10)  //查询id为10
	db.Where("name=?", "kodt_2").First(&user) // select * from users where name = 'kodt_2' limit 1;
	fmt.Printf("user:%v\n", user)
	// 可以传结构体进去查询
	// NOT查询 OR查询
	db.FirstOrInit(&user, User{Name: "non_existing"}) // 查不到创建
}
