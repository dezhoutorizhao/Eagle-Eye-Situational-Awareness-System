package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"utils"
)

type User struct {
	Id int
	UserName string
	Password string
	Email string
}

var (
	id int
	username string
	password string
	email string
)

func (user *User) GetUserByUsername(un string) (*User,error) {
	// sql语句
	sqlStr := "SELECT id, username, password, email FROM user_login.users WHERE username = ?";
	// QueryRow执行一次查询，并期望返回最多一行结果，即row
	fmt.Println(sqlStr)
	rows, err := utils.Db.Query(sqlStr, un)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &User{
			Id: id,
			UserName: username,
			Password: password,
			Email: email,
		}
		fmt.Println(u)
		return u, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

func checkLogin(username string, password string) int {
	// user是从前端接收的数据
	user := &User{}
	// u是根据username查询到的数据
	u, err := user.GetUserByUsername(username)
	fmt.Println(username,password)
	fmt.Println(u)
	if err != nil {
		fmt.Println(err)
		return 500 // Internal Server Error
	}
	if u == nil {
		fmt.Println("404")
		return 404 // Not Found
	}
	if u.Password == password {
		fmt.Println("login successfully")
		return 200 // OK
	} else {
		fmt.Println("401")
		return 401 // Unauthorized
	}
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	statusCode := checkLogin(username, password)
	if statusCode == 200 {
		// Login successful, redirect to dashboard or set cookies
	} else if statusCode == 401 {
		// Invalid credentials, display error message
	} else {
		// Other error, display appropriate message
	}
}