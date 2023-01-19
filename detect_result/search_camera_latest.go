package detect_result

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Search_camera(search_camera *gin.Context) {
	body, _ := search_camera.GetRawData()
	type Search struct {
		Number string `json:"number"`
	}
	var search Search
	err := json.Unmarshal(body, &search)
	if err != nil {
		fmt.Println(err.Error())
	}

	var camera Camera
	DB.Take(&camera, "number = ?", search.Number)
	var search_camera_return Camera
	search_camera_return.Number = camera.Number
	search_camera_return.Position = camera.Position
	search_camera_return.Task = camera.Task
	search_camera_return.Rtsp = camera.Rtsp
	search_camera_return.Remarks = camera.Remarks
	search_camera_return.Probability = camera.Probability
	search_camera_return.Framerate = camera.Framerate
	search_camera_return.Frameratetest = camera.Frameratetest
	search_camera.JSON(200, search_camera_return)
}
