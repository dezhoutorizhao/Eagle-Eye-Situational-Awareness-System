package db

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Modify_user(modify_user *gin.Context) {
	body, _ := modify_user.GetRawData()
	contentType := modify_user.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		var u User
		err := json.Unmarshal(body, &u)
		mod_sql := "UPDATE user_login.users set username = ?,password = ?,email = ?,role = ?,school_id = ? where id = ?"
		inStmt, err := Db.Prepare(mod_sql)
		if err != nil {
			fmt.Println("预编译出现异常", err)
			return
		}
		_, err2 := inStmt.Exec(u.UserName, u.Password, u.Email, u.Role, u.School_id, u.Id)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
		}

	}
}
