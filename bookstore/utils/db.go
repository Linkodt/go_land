package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:n246a369.@tcp(localhost:3306)/bookstore0612")
	if err != nil{
		panic(err.Error())
	}
}