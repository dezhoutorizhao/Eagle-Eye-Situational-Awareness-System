package detection

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		var Task []string
		var Rtsp string
		var Remarks string
		var Probability string
		var Framerate string
		var Frameratetest string

		err = rows.Scan(&Id, &Number, &Position, &Task, &Rtsp, &Remarks, &Probability, &Framerate, &Frameratetest)
		if err != nil {
			fmt.Println(err.Error())
		}

		camera_to_front := Camera{
			Id:            Id,
			Number:        Number,
			Position:      Position,
			Task:          Task,
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
