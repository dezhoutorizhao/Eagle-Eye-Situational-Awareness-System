package detection

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Delete_camera_struct struct {
	Id int `json:"id"`
}

func Delete_camera(delete_camera *gin.Context) {
	body, _ := delete_camera.GetRawData()
	var camera Delete_camera_struct
	err := json.Unmarshal(body, &camera)
	if err != nil {
		fmt.Println(err.Error())
	}

	mod_sql := "DELETE FROM detection.cameras WHERE id = ?"
	_, err = Db.Query(mod_sql, camera.Id)
	if err != nil {
		fmt.Println(err)
	}
}
