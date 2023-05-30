package db

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Review struct {
	Review_un        string
	Review_pw        string
	Review_em        string
	Review_role      string
	Review_school_id string
}

type Get_information struct {
	Get_un        string `json:"review_un"`
	Get_pw        string `json:"review_pw"`
	Get_em        string `json:"review_em"`
	Get_role      string `json:"review_role"`
	Get_school_id string `json:"review_school_id"`
	Get_whether   bool   `json:"review_whether"`
}

func Review_func(review_gin *gin.Context) {
	rows, err := Db.Query("SELECT * FROM user_login.register_users")
	// 用于存放Review结构体的切片
	reviews := make([]Review, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		var Review_un string
		var Review_pw string
		var Review_em string
		var Review_role string
		var Review_school_id string

		err = rows.Scan(&Review_un, &Review_pw, &Review_em, &Review_school_id, &Review_role)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		review := Review{
			Review_un:        Review_un,
			Review_pw:        Review_pw,
			Review_em:        Review_em,
			Review_role:      Review_role,
			Review_school_id: Review_school_id,
		}

		reviews = append(reviews, review)
		//fmt.Println(Review_em)

	}
	review_gin.JSON(200, reviews)
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
		if user.Get_whether == false {
			fmt.Println("不允许添加")
			delete_sql := "DELETE FROM user_login.register_users WHERE number = ?"
			println("number是", user.Get_school_id)
			_, err = Db.Query(delete_sql, user.Get_school_id)
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		if user.Get_role == "管理员" {
			sql_str := "insert into user_login.users values(?,?,?,?,?,?)"
			inStmt, err := Db.Prepare(sql_str)
			if err != nil {
				fmt.Println("预编译出现异常", err)
			}
			_, err2 := inStmt.Exec(nil, user.Get_un, user.Get_pw, user.Get_em, user.Get_role, user.Get_school_id)
			if err2 != nil {
				fmt.Println("执行出现异常", err2)
			}
		} else {
			sql_str := "insert into user_login.users values(?,?,?,?,?,?)"
			inStmt, err := Db.Prepare(sql_str)
			if err != nil {
				fmt.Println("预编译出现异常", err)
			}
			_, err2 := inStmt.Exec(nil, user.Get_un, user.Get_pw, user.Get_em, user.Get_role, user.Get_school_id)
			if err2 != nil {
				fmt.Println("执行出现异常", err2)
			}
		}
		// 添加完后删除
		delete_sql := "DELETE FROM user_login.register_users WHERE number = ?"
		println("number是", user.Get_school_id)
		_, err = Db.Query(delete_sql, user.Get_school_id)
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
