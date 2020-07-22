package main

// 自定义模型
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// 定义模型
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 // 零值类型
	birthday     *time.Time    // 改列名 `gorm:"column:birthday_time"`
	Email        string        `gorm:"type:varchar(100);unique_index"`
	Role         string        `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string       `gorm:"unique;not null"` // 设置会员号 唯一并且不为空
	Num          int           `gorm:"AUTO_INCREMENT"`  // 设置num为自增类型
	Address      string        `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int           `gorm:"-"`               // 忽略本字段

}

func (User) TableName() string { // 手动设置表名
	return "mono"
}

func main() {
	// 链接MySQL数据库
	db, err := gorm.Open("mysql", "root:n246a369.@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)  // 禁用复数  users-->user
	db.AutoMigrate(&User{}) // 自动创建表 自动迁移

	// 使用User结构体创建一个名叫xiaowangzi的表
	//db.Table("xiaowangzi").CreateTable(&User{})
	// 下面是创建表名时,修改创建没有被指定的表的表名
	//gorm.DefaultTableNameHandler = func(db *gorm.DB,defaultTableName string)string{
	//	return "SMS_" + defaultTableName
	//}
}
