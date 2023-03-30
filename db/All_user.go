package db

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
)

type All_user_struct struct {
	Id        int    `json:"user_id"`
	Un        string `json:"user_un"`
	Pw        string `json:"user_pw"`
	Em        string `json:"user_em"`
	Role      string `json:"user_role"`
	School_id string `json:"user_school_id"`
}

//func All_users(all_user *gin.Context) {
//
//	rows, _ := Db.Query("SELECT * FROM user_login.users")
//
//	all_users_to_front := make([]All_user_struct, 0)
//
//	for rows.Next() {
//		//var Id int
//		//var Un string
//		//var Pw string
//		//var Em string
//		//var Role string
//		//var School_id string
//
//		err = rows.Scan(&id, &username, &password, &email, &role, &school_id)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		user_to_front := All_user_struct{
//			Id:        id,
//			Un:        username,
//			Pw:        password,
//			Em:        email,
//			Role:      role,
//			School_id: school_id,
//		}
//
//		all_users_to_front = append(all_users_to_front, user_to_front)
//	}
//	all_user.JSON(200, all_users_to_front)
//}

func All_users(all_user *gin.Context) {
	rows, _ := Db.Query("SELECT * FROM user_login.users")
	all_users_to_front := make([]User, 0)
	for rows.Next() {
		_ = rows.Scan(&id, &username, &password, &email, &role, &school_id)
		u := User{
			Id:        id,
			UserName:  username,
			Password:  password,
			Email:     email,
			Role:      role,
			School_id: school_id,
		}
		all_users_to_front = append(all_users_to_front, u)
		//println(u)
	}
	all_user.JSON(200, all_users_to_front)
}
