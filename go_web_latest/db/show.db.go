package db

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Show(show *gin.Context) {
	type School_id struct {
		School_id string `json:"school_id"`
	}
	body, _ := show.GetRawData()
	contentType := show.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		var user School_id
		err := json.Unmarshal(body, &user)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(user)
		// 搜索sql语句
		sqlStatement := "SELECT * FROM user_login.users WHERE school_id = ?"
		// 执行
		rows, err := Db.Query(sqlStatement, user.School_id)
		fmt.Println(user.School_id)
		fmt.Println(rows)
		fmt.Println("yes")
		//result, err := json.Marshal(rows)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}

		var school_name string
		for rows.Next() {
			err := rows.Scan(nil, nil, nil, nil, nil, &school_name)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println(school_name)
		show.String(200, school_name)
	}
}
