package detect_result

//
//import (
//	_ "database/sql"
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//)
//
//type Mod_log_front struct {
//	Id int `json:"id"`
//}
//
//func Modify_logs_review(modify_logs *gin.Context) {
//	body, _ := modify_logs.GetRawData()
//	fmt.Println(body)
//	contentType := modify_logs.GetHeader("Content-Type")
//	switch contentType {
//	case "application/json":
//		var mod_log Mod_log_front
//		err := json.Unmarshal(body, &mod_log)
//		mod_sql := "UPDATE detection.results set review = ?"
//		inStmt, err := DB.Prepare(mod_sql)
//		if err != nil {
//			fmt.Println("预编译出现异常", err)
//			return
//		}
//
//		_, err2 := inStmt.Exec(mod_log.Id)
//		if err2 != nil {
//			fmt.Println("执行出现异常", err2)
//		}
//
//	}
//
//}

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Modify_logs_review(modify_logs *gin.Context) {
	body, _ := modify_logs.GetRawData()
	fmt.Println(body)
	contentType := modify_logs.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		type Mod_log_front struct {
			Id int `json:"id"`
		}

		var mod_log Mod_log_front
		err := json.Unmarshal(body, &mod_log)
		mod_sql := "UPDATE detection.results set review = ? WHERE id = ?"
		inStmt, err := Db_sql.Prepare(mod_sql)
		if err != nil {
			fmt.Println("预编译出现异常", err)
			return
		}
		_, err2 := inStmt.Exec(1, mod_log.Id)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
		}
	}
}
