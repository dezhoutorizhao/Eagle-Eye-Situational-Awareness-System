package detection

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func All_camera(all_camera *gin.Context) {

	rows, err := Db.Query("SELECT * FROM detection.cameras")

	all_cameras_to_front := make([]Camera, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var Id int
		var Number string
		var Position string
		var Task string
		var Rtsp string
		var Remarks string
		var Probability string
		var Framerate string
		var Frameratetest string
		var Task_to_front []string

		err = rows.Scan(&Id, &Number, &Position, &Task, &Rtsp, &Remarks, &Probability, &Framerate, &Frameratetest)
		if err != nil {
			fmt.Println(err.Error())
		}

		if strings.Contains(Task, "1") {
			Task_to_front = append(Task_to_front, "火灾")
		}
		if strings.Contains(Task, "2") {
			Task_to_front = append(Task_to_front, "吸烟")
		}
		if strings.Contains(Task, "3") {
			Task_to_front = append(Task_to_front, "栏杆")
		}
		if strings.Contains(Task, "4") {
			Task_to_front = append(Task_to_front, "挥手")
		}
		if strings.Contains(Task, "5") {
			Task_to_front = append(Task_to_front, "溺水")
		}
		if strings.Contains(Task, "6") {
			Task_to_front = append(Task_to_front, "摔倒")
		}

		camera_to_front := Camera{
			Id:            Id,
			Number:        Number,
			Position:      Position,
			Task:          Task_to_front,
			Rtsp:          Rtsp,
			Remarks:       Remarks,
			Probability:   Probability,
			Framerate:     Framerate,
			Frameratetest: Frameratetest,
		}

		all_cameras_to_front = append(all_cameras_to_front, camera_to_front)

	}

	all_camera.JSON(200, all_cameras_to_front)
}
