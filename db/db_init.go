package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)


func init() {
	Db,err = sql.Open("mysql","root:20030729a@tcp(localhost:3306)/user_login")
	if Db != nil {
		fmt.Println("Db is not nil")
	}
	if err != nil {
		fmt.Println(err)
	}
	Db.SetConnMaxLifetime(10)
	Db.SetMaxIdleConns(5)
	if err := Db.Ping() ; err != nil {
		fmt.Println("open database fail")
		return
	} else {
		println("yes")
	}
}
