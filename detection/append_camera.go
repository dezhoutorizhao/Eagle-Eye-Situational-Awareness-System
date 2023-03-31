package detection

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Camera struct {
	Id            int      `json:"id"`
	Number        string   `json:"number"`
	Position      string   `json:"position"`
	Task          []string `json:"task"`
	Rtsp          string   `json:"rtsp"`
	Remarks       string   `json:"remarks"`
	Probability   string   `json:"probability"`
	Framerate     string   `json:"framerate"`
	Frameratetest string   `json:"frameratetest"`
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// if_stop变量用来控制是否重启检测脚本
var If_stop bool = false

func Add_camera(add_camera *gin.Context) {

	body, _ := add_camera.GetRawData()
	fmt.Println(body)
	contentType := add_camera.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		// 读取部分
		var camera Camera
		err := json.Unmarshal(body, &camera)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(camera)
		// 添加部分
		sql_str := "insert into detection.cameras values(?,?,?,?,?,?,?,?,?)"
		inStmt, err := Db.Prepare(sql_str)
		if err != nil {
			fmt.Println("预编译出现异常", err)
			return
		}
		//fmt.Println(camera.Task)
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
		if Contains(camera.Task, "积水") {
			task_str += "7"
		}

		_, err2 := inStmt.Exec(nil, camera.Number, camera.Position, task_str, camera.Rtsp, camera.Remarks, camera.Probability, camera.Framerate, camera.Frameratetest)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
			return
		}
		If_stop = true
		fmt.Println("执行结束")
		If_stop = false
	}
}
