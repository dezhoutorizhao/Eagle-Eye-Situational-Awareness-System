package db


//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//)
//
//var (
//	re_username string
//	re_password string
//	re_email string
//)
//
//// 接收从前端返回过来的数据部分
//func AddUser_front(c *gin.Context) {
//	re_username = c.PostForm("username")
//	re_password = c.PostForm("password")
//	re_email = c.PostForm("email")
//	fmt.Println(re_username,re_password,re_email)
//	u := &User{re_username,re_password,re_email}
//	u.AddUser()
//}
//
//// AddUser 添加用户的方法一
//func (user *User) AddUser() error {
////写sql语句
//sqlStr := "insert into users(username,password,email) values(?,?,?)"
////预编译
//fmt.Println(sqlStr)
//inStmt , err := Db.Prepare(sqlStr) //预编译得到的是inStmt,通过操作inStmt得到不同的结果
//if err != nil {
//fmt.Println("预编译出现异常",err)
//return err
//}
////3.执行
//fmt.Println(inStmt)
//_,err2 := inStmt.Exec(re_username,re_password,re_email)
//if err2 != nil {
//fmt.Println("执行出现异常",err2)
//return err2
//}
//return nil
//}