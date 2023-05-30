package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		panic("MYSQL_HOST environment variable not set")
	}
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		panic("MYSQL_PORT environment variable not set")
	}
	Db, err = sql.Open("mysql", "root:20030729a@tcp("+host+":"+port+")/user_login")
	if Db != nil {
		fmt.Println("Db is not nil")
	}
	if err != nil {
		fmt.Println(err)
	}
	Db.SetConnMaxLifetime(10)
	Db.SetMaxIdleConns(5)
	if err := Db.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	} else {
		println("yes")
	}
}
func BindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}
