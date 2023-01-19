package detect_result

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
	"strconv"
	"strings"
)

func StringToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func StringToFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		// handle error
		return 0
	}
	return float32(f)
}

func Get_camera() {
	var cameraList []Camera

	//返回查询的结果条数
	count := DB.Find(&cameraList).RowsAffected
	fmt.Println(count)

	// 从数据库中获取的参数
	var rtsp_data string
	var vid_stride int
	var threshold float32
	var detect_task string

	for _, camera := range cameraList {
		// 获取参数
		rtsp_data = camera.Rtsp
		vid_stride = StringToInt(camera.Framerate)
		threshold = StringToFloat32(camera.Frameratetest)
		// 获取task执行参数，并进行调用
		detect_task = camera.Task
		fire := strings.Contains(detect_task, "1")
		smoke := strings.Contains(detect_task, "2")
		railing := strings.Contains(detect_task, "3")
		wave := strings.Contains(detect_task, "4")
		drown := strings.Contains(detect_task, "5")
		fall := strings.Contains(detect_task, "6")
		if fire != false {
			Run_python_fire(rtsp_data, vid_stride, threshold, nil)
		}
		if smoke != false {
			Run_python_smoke(rtsp_data, vid_stride, threshold, nil)
		}
		if railing != false {
			Run_python_railing(rtsp_data, vid_stride, threshold, nil)
		}
		if wave != false {
			Run_python_wave(rtsp_data, vid_stride, threshold, nil)
		}
		if drown != false {
			Run_python_drown(rtsp_data, vid_stride, threshold, nil)
		}
		if fall != false {
			Run_python_fall(rtsp_data, vid_stride, threshold, nil)
		}
	}
}

func Run_python_fire(data string, vid_stride int, threshold float32, fire *gin.Context) {
	cmd := exec.Command("python", "-c", "import fire;fire.fire()")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fire.JSON(200, stdout)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(stdout, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 := Results{
		Photo: result.Photo,
		Video: result.Video,
		Rate:  result.Rate,
		Task:  result.Task,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)
}

func Run_python_smoke(data string, vid_stride int, threshold float32, fire *gin.Context) {
	cmd := exec.Command("python", "-c", "import fire;fire.fire()")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fire.JSON(200, stdout)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(stdout, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 := Results{
		Photo: result.Photo,
		Video: result.Video,
		Rate:  result.Rate,
		Task:  result.Task,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)
}

func Run_python_railing(data string, vid_stride int, threshold float32, fire *gin.Context) {
	cmd := exec.Command("python", "-c", "import fire;fire.fire()")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fire.JSON(200, stdout)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(stdout, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 := Results{
		Photo: result.Photo,
		Video: result.Video,
		Rate:  result.Rate,
		Task:  result.Task,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)
}

func Run_python_wave(data string, vid_stride int, threshold float32, fire *gin.Context) {
	cmd := exec.Command("python", "-c", "import fire;fire.fire()")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fire.JSON(200, stdout)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(stdout, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 := Results{
		Photo: result.Photo,
		Video: result.Video,
		Rate:  result.Rate,
		Task:  result.Task,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)
}

func Run_python_drown(data string, vid_stride int, threshold float32, fire *gin.Context) {
	cmd := exec.Command("python", "-c", "import fire;fire.fire()")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fire.JSON(200, stdout)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(stdout, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 := Results{
		Photo: result.Photo,
		Video: result.Video,
		Rate:  result.Rate,
		Task:  result.Task,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)
}

func Run_python_fall(data string, vid_stride int, threshold float32, fire *gin.Context) {
	cmd := exec.Command("python", "-c", "import fire;fire.fire()")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fire.JSON(200, stdout)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(stdout, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 := Results{
		Photo: result.Photo,
		Video: result.Video,
		Rate:  result.Rate,
		Task:  result.Task,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)
}
