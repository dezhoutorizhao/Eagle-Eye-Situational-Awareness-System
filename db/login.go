package db

import (
	_ "context"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id        int    `json:"id"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	School_id string `json:"school_id"`
}

type User_login struct {
	Id            int    `json:"id"`
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	School_id     string `json:"school_id"`
	Whether_login string `json:"whether_login"`
}

var (
	id           int
	username     string
	password     string
	email        string
	login_status int
	role         string
	school_id    string
)

func (user *User) GetUserByUsername(un string) (*User, error) {
	// sql语句
	sqlStr := "SELECT id, username, password, email,role,school_id FROM user_login.users WHERE username = ?"
	// QueryRow执行一次查询，并期望返回最多一行结果，即row
	fmt.Println(sqlStr)
	if Db == nil {
		return nil, fmt.Errorf("Db is nil")
	}
	rows, err := Db.Query(sqlStr, un)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &username, &password, &email, &role, &school_id)
		if err != nil {
			return nil, err
		}
		u := &User{
			Id:        id,
			UserName:  username,
			Password:  password,
			Email:     email,
			Role:      role,
			School_id: school_id,
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
	fmt.Println(username, password)
	fmt.Println(u)
	if err != nil {
		fmt.Println(err)
		login_status = 0
		return 500 // Internal Server Error
	}
	if u == nil {
		fmt.Println("404")
		login_status = 0
		return 404 // Not Found
	}
	if u.Password == password {
		fmt.Println("login successfully")
		//log_status := If_Success{"yes"}
		//c.JSON(200,log_status)
		login_status = 200
		return 200 // OK
	} else {
		fmt.Println("401")
		login_status = 0
		return 401 // Unauthorized
	}
}

func Login(c *gin.Context) {
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	body, _ := c.GetRawData()
	type login struct {
		UName string `json:"username"`
		PWord string `json:"password"`
	}

	var user login
	err := json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
	}

	statusCode := checkLogin(user.UName, user.PWord)

	fmt.Println(statusCode)

	if statusCode == 200 {
		// Login successful, redirect to dashboard or set cookies
	} else if statusCode == 401 {
		// Invalid credentials, display error message
	} else {
		// Other error, display appropriate message
	}

}

func If_success(c *gin.Context) {
	if login_status == 200 {
		//c.String(200, "success")
		user_whole := User_login{
			Id:            id,
			UserName:      username,
			Password:      password,
			Email:         email,
			Role:          role,
			School_id:     school_id,
			Whether_login: "success",
		}

		c.JSON(200, user_whole)
	}
}
