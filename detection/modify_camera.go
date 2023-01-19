package detection

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

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
		sql_str := "UPDATE detection.cameras set position = ?,task = ?,rtsp = ?,remarks = ?,probability= ?,framerate = ?,frameratetest = ? where number = ?"
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
		_, err2 := inStmt.Exec(camera.Position, task_str, camera.Rtsp, camera.Remarks, camera.Probability, camera.Framerate, camera.Frameratetest, camera.Number)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
			return
		}
		fmt.Println("执行结束")
	}
}
