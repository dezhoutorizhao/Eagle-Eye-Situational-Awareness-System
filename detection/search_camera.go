package detection

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Serach_camera(search_camera *gin.Context) {
	// 获取要查询的数据
	//search_number := search_camera.PostForm("number")
	//
	//// 写sql语句
	//sql_str := "select * FROM detection.camera where number = ?"
	//inStmt, err := Db.Prepare(sql_str)
	//if err != nil {
	//	fmt.Println("预编译出现异常", err)
	//	return
	//}
	//rows, err2 := inStmt.Exec(search_number)
	//if err2 != nil {
	//	fmt.Println("执行出现异常", err2)
	//	return
	//}
	//fmt.Println(rows)

	number := search_camera.Param("number")

	// 查询数据库
	rows, err := Db.Query("select * FROM detection.camera where number = ?", number)
	if err != nil {
		fmt.Println("查询失败！")
		return
	}
	var camera Camera
	for rows.Next() {
		err := rows.Scan(&camera.Number, &camera.Position, &camera.Task, &camera.Rtsp, &camera.Remarks, &camera.Probability, &camera.Framerate, &camera.Frameratetest)
		if err != nil {
			fmt.Println("读取查询结果失败！")
			return
		}
	}
	data, err := json.Marshal(camera)
	if err != nil {
		fmt.Println("转换json格式失败！")
		return
	}
	search_camera.Data(http.StatusOK, "application/json", data)
}
