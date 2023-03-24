package detect_result

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

func getDatetime() sql.NullString {
	now := time.Now()
	datetime := now.Format("2006-01-02 15:04:05")
	return sql.NullString{
		String: datetime,
		Valid:  true,
	}
}

func Detect_process(get_camera *gin.Context) {

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

		Judge_fire(fire, rtsp_data, vid_stride, threshold, where_loc, get_camera)
		Judge_smoke(smoke, rtsp_data, vid_stride, threshold, where_loc, get_camera)
		Judge_railing(railing, rtsp_data, vid_stride, threshold, where_loc, get_camera)
		Judge_wave(wave, rtsp_data, vid_stride, threshold, where_loc, get_camera)
		Judge_drown(drown, rtsp_data, vid_stride, threshold, where_loc, get_camera)
		Judge_fall(fall, rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
	// 如果正在修改,If_stop变量的值就更改为true,此时关停，当再次访问路由时重新启动
	//if detection.If_stop == true {
	//	return false
	//}
	//return true
}

func Judge_fire(fire bool, rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {
	//fire
	if fire != false {
		timeout := time.After(5 * time.Second) // 设置超时时间为 5 秒
		ticker := time.Tick(300 * time.Millisecond)
		for {
			select {
			case <-timeout:
				fmt.Println("超时")
				time.Sleep(300 * time.Millisecond)
				Run_python_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
			case <-ticker:
				Run_python_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
			}
		}
		fmt.Println(rtsp_data, " ", vid_stride, " ", threshold, " ", where_loc)
	}

	////fire
	//if fire != false {
	//	timeout := time.After(5 * time.Second) // 设置超时时间为 5 秒
	//	ticker := time.Tick(300 * time.Millisecond)
	//	for {
	//		select {
	//		case <-timeout:
	//			fmt.Println("超时")
	//			return
	//		case <-ticker:
	//			Run_python_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	//		}
	//	}
	//	fmt.Println(rtsp_data, " ", vid_stride, " ", threshold, " ", where_loc)
	//}
}

func Judge_smoke(smoke bool, rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {
	//smoke
	if smoke != false {
		timeout := time.After(5 * time.Second) // 设置超时时间为 5 秒
		ticker := time.Tick(300 * time.Millisecond)
		for {
			select {
			case <-timeout:
				fmt.Println("超时")
				Judge_smoke(smoke, rtsp_data, vid_stride, threshold, where_loc, get_camera)
			case <-ticker:
				Run_python_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
			}
		}
		fmt.Println(rtsp_data, " ", vid_stride, " ", threshold, " ", where_loc)
	}
}

func Judge_railing(railing bool, rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {
	//railing
	if railing != false {
		timeout := time.After(5 * time.Second) // 设置超时时间为 5 秒
		ticker := time.Tick(300 * time.Millisecond)
		for {
			select {
			case <-timeout:
				fmt.Println("超时")
				Judge_railing(railing, rtsp_data, vid_stride, threshold, where_loc, get_camera)
			case <-ticker:
				Run_python_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
			}
		}
		fmt.Println(rtsp_data, " ", vid_stride, " ", threshold, " ", where_loc)
	}
}

func Judge_wave(wave bool, rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {
	//wave
	if wave != false {
		timeout := time.After(5 * time.Second) // 设置超时时间为 5 秒
		ticker := time.Tick(300 * time.Millisecond)
		for {
			select {
			case <-timeout:
				fmt.Println("超时")
				time.Sleep(300 * time.Millisecond)
				Judge_wave(wave, rtsp_data, vid_stride, threshold, where_loc, get_camera)
			case <-ticker:
				Run_python_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
			}
		}

	}
}

func Judge_drown(drown bool, rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {
	//drown
	if drown != false {
		timeout := time.After(5 * time.Second) // 设置超时时间为 5 秒
		ticker := time.Tick(300 * time.Millisecond)
		for {
			select {
			case <-timeout:
				fmt.Println("超时")
				Judge_drown(drown, rtsp_data, vid_stride, threshold, where_loc, get_camera)
			case <-ticker:
				Run_python_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
			}
		}
		fmt.Println(rtsp_data, " ", vid_stride, " ", threshold, " ", where_loc)
	}
}

func Judge_fall(fall bool, rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {
	if fall != false {
		timeout := time.After(5 * time.Second) // 设置超时时间为 5 秒
		ticker := time.Tick(300 * time.Millisecond)
		for {
			select {
			case <-timeout:
				fmt.Println("超时")
				Judge_fall(fall, rtsp_data, vid_stride, threshold, where_loc, get_camera)
			case <-ticker:
				Run_python_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
			}
		}
		fmt.Println(rtsp_data, " ", vid_stride, " ", threshold, " ", where_loc)
	}
}

//以上代码中，timeout 是一个 time.After() 函数返回的 chan time.Time 类型的 channel，当超过设定时间后会从这个 channel 中接收到一个时间值，表示已经超时了。在 for 循环中，我们通过 select 语句来监听 timeout 和 ticker 这两个 channel，如果 timeout 先接收到了值，那么就退出循环，停止调用 Run_python_wave() 函数。如果 ticker 先接收到了值，那么就调用 Run_python_wave() 函数，然后等待下一次循环。这样就可以实现在超过一定时间后停止函数的调用。

func Run_python_fire(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {

	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "./Files/Fire", "/dtest.py")
	fmt.Println(pythonFile)

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_threshold, "--interval", temp_vid_stride, "--location", where_loc, "--task", "fire")
	fmt.Println(cmd)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(out))

	//from
	//var algorithm_result map[string]interface{}
	//
	//fire.String(200, "1")
	//
	////解析并添加到数据库中
	//trans_err := json.Unmarshal(out, &algorithm_result)
	//if err != nil {
	//	fmt.Println(trans_err)
	//}

	//temp_photo := algorithm_result["Photo"]
	//temp_location := algorithm_result["Location"]
	//temp_rate := algorithm_result["Rate"]
	//temp_task := algorithm_result["Task"]
	//temp_Review := sql.NullString{
	//	String: "false",
	//	Valid:  true,
	//}

	//end
	//temp_Review := 0
	//
	//// 如果检测结果为空，则不去添加进数据库
	//if algorithm_result["Photo"] == nil {
	//	fmt.Println("未检测到结果")
	//	return
	//} else {
	//	this_time := getDatetime()
	//	s1 := Add_to_database{
	//		Photo:    algorithm_result["Photo"].(sql.NullString),
	//		Rate:     algorithm_result["Rate"].(sql.NullFloat64),
	//		Task:     algorithm_result["Task"].(sql.NullString),
	//		Location: algorithm_result["Location"].(sql.NullString),
	//		Time:     this_time,
	//		Review:   temp_Review,
	//	}
	//	create_err := DB.Create(&s1).Error
	//	fmt.Println(create_err)

	// 传递给公众号
	//To_weixin(result)
	//}

}

func Run_python_smoke(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "./Files/Fire", "/dtest.py")

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_vid_stride, "--interval", temp_threshold, "--location", where_loc, "--task", "fire")

	out, err := cmd.Output()
	//fmt.Println(out.String())

	var algorithm_result map[string]interface{}

	fire.String(200, string(out))

	//解析并添加到数据库中
	trans_err := json.Unmarshal(out, &algorithm_result)
	if err != nil {
		fmt.Println(trans_err)
	}
	temp_photo := algorithm_result["Photo"]
	temp_location := algorithm_result["Location"]
	temp_rate := algorithm_result["Rate"]
	temp_task := algorithm_result["Task"]
	//temp_Review := sql.NullString{
	//	String: "false",
	//	Valid:  true,
	//}
	temp_Review := 0

	// 如果检测结果为空，则不去添加进数据库
	if temp_photo == nil {
		fmt.Println("未检测到结果")
		return
	} else {
		this_time := getDatetime()
		s1 := Add_to_database{
			Photo:    temp_photo.(sql.NullString),
			Rate:     temp_rate.(sql.NullFloat64),
			Task:     temp_task.(sql.NullString),
			Location: temp_location.(sql.NullString),
			Time:     this_time,
			Review:   temp_Review,
		}
		create_err := DB.Create(&s1).Error
		fmt.Println(create_err)

		// 传递给公众号
		//To_weixin(result)
	}
}

func Run_python_railing(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "./Files/Fire", "/dtest.py")

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_vid_stride, "--interval", temp_threshold, "--location", where_loc, "--task", "fire")

	out, err := cmd.Output()
	//fmt.Println(out.String())

	var algorithm_result map[string]interface{}

	fire.String(200, string(out))

	//解析并添加到数据库中
	trans_err := json.Unmarshal(out, &algorithm_result)
	if err != nil {
		fmt.Println(trans_err)
	}
	temp_photo := algorithm_result["Photo"]
	temp_location := algorithm_result["Location"]
	temp_rate := algorithm_result["Rate"]
	temp_task := algorithm_result["Task"]
	//temp_Review := sql.NullString{
	//	String: "false",
	//	Valid:  true,
	//}
	temp_Review := 0

	// 如果检测结果为空，则不去添加进数据库
	if temp_photo == nil {
		fmt.Println("未检测到结果")
		return
	} else {
		this_time := getDatetime()
		s1 := Add_to_database{
			Photo:    temp_photo.(sql.NullString),
			Rate:     temp_rate.(sql.NullFloat64),
			Task:     temp_task.(sql.NullString),
			Location: temp_location.(sql.NullString),
			Time:     this_time,
			Review:   temp_Review,
		}
		create_err := DB.Create(&s1).Error
		fmt.Println(create_err)

		// 传递给公众号
		//To_weixin(result)
	}
}

func Run_python_wave(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "./Files/Fire", "/dtest.py")

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_vid_stride, "--interval", temp_threshold, "--location", where_loc, "--task", "fire")

	out, err := cmd.Output()
	//fmt.Println(out.String())

	var algorithm_result map[string]interface{}

	fire.String(200, string(out))

	//解析并添加到数据库中
	trans_err := json.Unmarshal(out, &algorithm_result)
	if err != nil {
		fmt.Println(trans_err)
	}
	temp_photo := algorithm_result["Photo"]
	temp_location := algorithm_result["Location"]
	temp_rate := algorithm_result["Rate"]
	temp_task := algorithm_result["Task"]
	//temp_Review := sql.NullString{
	//	String: "false",
	//	Valid:  true,
	//}
	temp_Review := 0

	// 如果检测结果为空，则不去添加进数据库
	if temp_photo == nil {
		fmt.Println("未检测到结果")
		return
	} else {
		this_time := getDatetime()
		s1 := Add_to_database{
			Photo:    temp_photo.(sql.NullString),
			Rate:     temp_rate.(sql.NullFloat64),
			Task:     temp_task.(sql.NullString),
			Location: temp_location.(sql.NullString),
			Time:     this_time,
			Review:   temp_Review,
		}
		create_err := DB.Create(&s1).Error
		fmt.Println(create_err)

		// 传递给公众号
		//To_weixin(result)
	}
}

func Run_python_drown(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "./Files/Fire", "/dtest.py")

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_vid_stride, "--interval", temp_threshold, "--location", where_loc, "--task", "fire")

	out, err := cmd.Output()
	//fmt.Println(out.String())

	var algorithm_result map[string]interface{}

	fire.String(200, string(out))

	//解析并添加到数据库中
	trans_err := json.Unmarshal(out, &algorithm_result)
	if err != nil {
		fmt.Println(trans_err)
	}
	temp_photo := algorithm_result["Photo"]
	temp_location := algorithm_result["Location"]
	temp_rate := algorithm_result["Rate"]
	temp_task := algorithm_result["Task"]
	//temp_Review := sql.NullString{
	//	String: "false",
	//	Valid:  true,
	//}
	temp_Review := 0

	// 如果检测结果为空，则不去添加进数据库
	if temp_photo == nil {
		fmt.Println("未检测到结果")
		return
	} else {
		this_time := getDatetime()
		s1 := Add_to_database{
			Photo:    temp_photo.(sql.NullString),
			Rate:     temp_rate.(sql.NullFloat64),
			Task:     temp_task.(sql.NullString),
			Location: temp_location.(sql.NullString),
			Time:     this_time,
			Review:   temp_Review,
		}
		create_err := DB.Create(&s1).Error
		fmt.Println(create_err)

		// 传递给公众号
		//To_weixin(result)
	}
}

func Run_python_fall(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "./Files/Fire", "/dtest.py")

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_vid_stride, "--interval", temp_threshold, "--location", where_loc, "--task", "fire")

	out, err := cmd.Output()
	//fmt.Println(out.String())

	var algorithm_result map[string]interface{}

	fire.String(200, string(out))

	//解析并添加到数据库中
	trans_err := json.Unmarshal(out, &algorithm_result)
	if err != nil {
		fmt.Println(trans_err)
	}
	temp_photo := algorithm_result["Photo"]
	temp_location := algorithm_result["Location"]
	temp_rate := algorithm_result["Rate"]
	temp_task := algorithm_result["Task"]
	//temp_Review := sql.NullString{
	//	String: "false",
	//	Valid:  true,
	//}
	temp_Review := 0

	// 如果检测结果为空，则不去添加进数据库
	if temp_photo == nil {
		fmt.Println("未检测到结果")
		return
	} else {
		this_time := getDatetime()
		s1 := Add_to_database{
			Photo:    temp_photo.(sql.NullString),
			Rate:     temp_rate.(sql.NullFloat64),
			Task:     temp_task.(sql.NullString),
			Location: temp_location.(sql.NullString),
			Time:     this_time,
			Review:   temp_Review,
		}
		create_err := DB.Create(&s1).Error
		fmt.Println(create_err)

		// 传递给公众号
		//To_weixin(result)
	}
}
