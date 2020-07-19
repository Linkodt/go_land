package utils
// 链接数据库
import (
	_"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init(){
	Db,err = sql.Open("mysql", "root:n246a369.@tcp(localhost:3306)/books")
	if err!=nil{
		panic(err.Error())
	}
}


func main() {
	
}