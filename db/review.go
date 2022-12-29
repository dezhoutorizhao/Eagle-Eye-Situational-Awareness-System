package db

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Review struct {
	Review_un  string
	Review_pw  string
	Review_em  string
	Review_num string
}

type Get_information struct {
	Get_un   string `json:"review_un"`
	Get_pw   string `json:"review_pw"`
	Get_em   string `json:"review_em"`
	Get_role string `json:"review_role"`
}

func Review_func(review_gin *gin.Context) {
	rows, err := Db.Query("SELECT * FROM user_login.register_users")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		var Review_un string
		var Review_pw string
		var Review_em string
		var Review_num string

		err = rows.Scan(&Review_un, &Review_pw, &Review_em, &Review_num)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// 用于存放Review结构体的切片
		reviews := make([]Review, 0)
		review := Review{
			Review_un:  Review_un,
			Review_pw:  Review_pw,
			Review_em:  Review_em,
			Review_num: Review_num,
		}

		reviews = append(reviews, review)
		//fmt.Println(Review_em)
		review_gin.JSON(200, reviews)
	}
}

func Get_Review(get_review *gin.Context) {
	body, _ := get_review.GetRawData()
	contentType := get_review.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		var user Get_information
		err := json.Unmarshal(body, &user)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(user)
		if user.Get_role == "administrator" {
			sql_str := "insert into user_login.users values(?,?,?,?,?)"
			inStmt, err := Db.Prepare(sql_str)
			if err != nil {
				fmt.Println("预编译出现异常", err)
			}
			_, err2 := inStmt.Exec(nil, user.Get_un, user.Get_pw, user.Get_em, user.Get_role)
			if err2 != nil {
				fmt.Println("执行出现异常", err2)
			}
		} else {
			sql_str := "insert into user_login.users values(?,?,?,?,?)"
			inStmt, err := Db.Prepare(sql_str)
			if err != nil {
				fmt.Println("预编译出现异常", err)
			}
			_, err2 := inStmt.Exec(nil, user.Get_un, user.Get_pw, user.Get_em, user.Get_role)
			if err2 != nil {
				fmt.Println("执行出现异常", err2)
			}
		}
	}
}
