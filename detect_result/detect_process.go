package detect_result

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"strongwill.com/detection"
)

var s1 Results

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

func Get_camera(get_camera *gin.Context) bool {

	var cameraList []Camera

	//返回查询的结果条数
	count := DB.Find(&cameraList).RowsAffected
	fmt.Println(count)

	// 从数据库中获取的参数
	var rtsp_data string
	var vid_stride int     //间隔的帧数
	var threshold float32  // 阈值的置信度
	var detect_task string // fire之类的
	var where_loc string   // 位置

	for _, camera := range cameraList {
		// 获取参数
		rtsp_data = camera.Rtsp
		where_loc = camera.Position
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
			Run_python_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
		}
		if smoke != false {
			Run_python_smoke(rtsp_data, vid_stride, threshold, where_loc, get_camera)
		}
		if railing != false {
			Run_python_railing(rtsp_data, vid_stride, threshold, where_loc, get_camera)
		}
		if wave != false {
			Run_python_wave(rtsp_data, vid_stride, threshold, where_loc, get_camera)
		}
		if drown != false {
			Run_python_drown(rtsp_data, vid_stride, threshold, where_loc, get_camera)
		}
		if fall != false {
			Run_python_fall(rtsp_data, vid_stride, threshold, where_loc, get_camera)
		}
	}
	// 如果正在修改,If_stop变量的值就更改为true,此时关停，当再次访问路由时重新启动
	if detection.If_stop == true {
		return false
	}
	return true
}

func Run_python_fire(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {

	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "detect_cls.py")

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, data, temp_vid_stride, temp_threshold, where_loc, "fire")

	out, err := cmd.Output()
	//fmt.Println(out.String())
	fire.JSON(200, out)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(out, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 = Results{
		Photo:    result.Photo,
		Rate:     result.Rate,
		Task:     result.Task,
		Location: result.Location,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)

	// 传递给公众号
	To_weixin(result)
}

func Run_python_smoke(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "detect_cls.py")

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, data, temp_vid_stride, temp_threshold, where_loc, "fire")

	out, err := cmd.Output()
	//fmt.Println(out.String())
	fire.JSON(200, out)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(out, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 = Results{
		Photo:    result.Photo,
		Rate:     result.Rate,
		Task:     result.Task,
		Location: result.Location,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)
}

func Run_python_railing(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "detect_cls.py")

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, data, temp_vid_stride, temp_threshold, where_loc, "fire")

	out, err := cmd.Output()
	//fmt.Println(out.String())
	fire.JSON(200, out)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(out, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 = Results{
		Photo:    result.Photo,
		Rate:     result.Rate,
		Task:     result.Task,
		Location: result.Location,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)
}

func Run_python_wave(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "detect_cls.py")

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, data, temp_vid_stride, temp_threshold, where_loc, "fire")

	out, err := cmd.Output()
	//fmt.Println(out.String())
	fire.JSON(200, out)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(out, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 = Results{
		Photo:    result.Photo,
		Rate:     result.Rate,
		Task:     result.Task,
		Location: result.Location,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)
}

func Run_python_drown(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "detect_cls.py")

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, data, temp_vid_stride, temp_threshold, where_loc, "fire")

	out, err := cmd.Output()
	//fmt.Println(out.String())
	fire.JSON(200, out)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(out, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 = Results{
		Photo:    result.Photo,
		Rate:     result.Rate,
		Task:     result.Task,
		Location: result.Location,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)
}

func Run_python_fall(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "detect_cls.py")

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, data, temp_vid_stride, temp_threshold, where_loc, "fire")

	out, err := cmd.Output()
	//fmt.Println(out.String())
	fire.JSON(200, out)

	//解析并添加到数据库中
	var result Results
	trans_err := json.Unmarshal(out, &result)
	if err != nil {
		fmt.Println(trans_err)
	}
	s1 = Results{
		Photo:    result.Photo,
		Rate:     result.Rate,
		Task:     result.Task,
		Location: result.Location,
	}
	create_err := DB.Create(&s1).Error
	fmt.Println(create_err)
}
