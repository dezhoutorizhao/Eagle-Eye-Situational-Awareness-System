package detection

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:20030729a@tcp(localhost:3306)/detection")
	if Db != nil {
		fmt.Println("Db is not nil")
	}
	if err != nil {
		fmt.Println(err)
	}
	Db.SetConnMaxLifetime(10)
	Db.SetMaxIdleConns(5)
	if err := Db.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	} else {
		println("yes")
	}
}

func Modify_camera(modify_camera *gin.Context) {
	body, _ := modify_camera.GetRawData()
	fmt.Println(body)
	contentType := modify_camera.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		var camera Camera
		err := json.Unmarshal(body, &camera)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(camera)
		sql_str := "UPDATE detection.cameras set number = ?,position = ?,task = ?,rtsp = ?,remarks = ?,probability= ?,framerate = ?,frameratetest = ? where id = ?"
		inStmt, err := Db.Prepare(sql_str)
		if err != nil {
			fmt.Println("预编译出现异常", err)
			return
		}
		task_str := ""
		if Contains(camera.Task, "火灾") {
			task_str += "1"
		}
		if Contains(camera.Task, "吸烟") {
			task_str += "2"
		}
		if Contains(camera.Task, "栏杆") {
			task_str += "3"
		}
		if Contains(camera.Task, "挥手") {
			task_str += "4"
		}
		if Contains(camera.Task, "溺水") {
			task_str += "5"
		}
		if Contains(camera.Task, "摔倒") {
			task_str += "6"
		}
		_, err2 := inStmt.Exec(camera.Number, camera.Position, task_str, camera.Rtsp, camera.Remarks, camera.Probability, camera.Framerate, camera.Frameratetest, camera.Id)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
			return
		}
		If_stop = true
		fmt.Println("执行结束")
		If_stop = false
	}
}
