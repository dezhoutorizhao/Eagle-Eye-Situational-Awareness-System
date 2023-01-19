package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	re_username string
	re_password string
	re_email    string
	re_number   string
)

// re代表register
type Re_User struct {
	re_username string `json:"username"`
	re_password string `json:"password"`
	re_email    string `json:"email"`
	re_number   string `json:"school_id"`
}

// 接收从前端返回过来的数据部分
func AddUser_front(add_c *gin.Context) {
	re_username = add_c.PostForm("username")
	re_password = add_c.PostForm("password")
	re_email = add_c.PostForm("email")
	re_number = add_c.PostForm("school_id")
	fmt.Println(re_username, re_password, re_email)
	if len(re_username) == 0 || len(re_password) == 0 || len(re_email) == 0 || len(re_number) == 0 {
		add_c.String(200, "error\n")
		return
	}
	if len(re_username) == 0 {
		add_c.String(200, "username is null\n")
		return
	}
	if len(re_password) == 0 {
		add_c.String(200, "password is null\n")
		return
	}
	if len(re_email) == 0 {
		add_c.String(200, "email is null\n")
		return
	}
	if len(re_number) == 0 {
		add_c.String(200, "school_id is null\n")
		return
	}
	// re代表register
	u := &Re_User{re_username, re_password, re_email, re_number}
	u.AddUser()
}

// AddUser 添加用户的方法一
func (user *Re_User) AddUser() error {
	//写sql语句
	sqlStr := "insert into register_users(username,password,email,number) values(?,?,?,?)"
	//预编译
	fmt.Println(sqlStr)
	inStmt, err := Db.Prepare(sqlStr) //预编译得到的是inStmt,通过操作inStmt得到不同的结果
	if err != nil {
		fmt.Println("预编译出现异常", err)
		return err
	}
	//3.执行
	fmt.Println(inStmt)
	_, err2 := inStmt.Exec(re_username, re_password, re_email, re_number)
	if err2 != nil {
		fmt.Println("执行出现异常", err2)
		return err2
	}
	return nil
}