package db

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Delete_user_id struct {
	Id int `json:"id"`
}

func Delete_user(delete_user *gin.Context) {
	body, _ := delete_user.GetRawData()
	var user Delete_user_id
	err := json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
	}

	mod_sql := "DELETE FROM user_login.users WHERE id = ?"
	_, err = Db.Query(mod_sql, user.Id)
	if err != nil {
		fmt.Println(err)
	}
}
