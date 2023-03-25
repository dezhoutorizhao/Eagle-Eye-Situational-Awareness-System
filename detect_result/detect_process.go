package detect_result

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"sync"
	"time"
)

var s1 Results
var mutex sync.Mutex

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
	// 查询数据库，将结果赋值给cameraList
	DB.Find(&cameraList)
	// 返回查询的结果条数
	count := len(cameraList)
	fmt.Println(count)
	// 创建一个waitgroup，用于等待所有goroutine执行完毕
	var wg sync.WaitGroup
	// 遍历cameraList，为每个camera记录创建一个goroutine并调用Judge函数
	for i := 0; i < count; i++ {
		wg.Add(1)
		camera := cameraList[i]
		go func(cam Camera) {
			defer wg.Done()
			var rtsp_data string
			var vid_stride int     //间隔的帧数
			var threshold float32  // 阈值的置信度
			var detect_task string // fire之类的
			var where_loc string   // 位置
			// 获取参数
			rtsp_data = cam.Rtsp
			where_loc = cam.Position
			vid_stride = StringToInt(cam.Framerate)
			if vid_stride == 0 {
				vid_stride = 5 // 默认值
			}
			threshold = StringToFloat32(cam.Frameratetest)
			if threshold == 0 {
				threshold = 0.5 // 默认值
			}
			// 获取task执行参数，并进行调用
			detect_task = cam.Task
			fire := strings.Contains(detect_task, "1")
			smoke := strings.Contains(detect_task, "2")
			railing := strings.Contains(detect_task, "3")
			wave := strings.Contains(detect_task, "4")
			drown := strings.Contains(detect_task, "5")
			fall := strings.Contains(detect_task, "6")
			Judge(fire, smoke, railing, wave, drown, fall, rtsp_data, vid_stride, threshold, where_loc, get_camera)
		}(camera)
	}
	// 等待所有goroutine执行完毕
	wg.Wait()
}

func Judge(fire bool, smoke bool, railing bool, wave bool, drown bool, fall bool, rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {

	//在互斥锁保护下调用fire、smoke、railing、wave、drown、fall等函数

	if fire != false {
		Judge_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
	if smoke != false {
		Judge_smoke(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
	if railing != false {
		Judge_railing(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
	if wave != false {
		Judge_wave(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
	if drown != false {
		Judge_drown(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
	if fall != false {
		Judge_fall(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
	println("正在跑")
}

func Judge_fire(rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {

	res := Run_python_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	if res == true {
		Judge_fire(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
}

func Judge_smoke(rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {
	res := Run_python_smoke(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	if res == true {
		Judge_smoke(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
}

func Judge_railing(rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {
	res := Run_python_railing(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	if res == true {
		Judge_railing(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
}

func Judge_wave(rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {
	res := Run_python_wave(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	if res == true {
		Judge_wave(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
}

func Judge_drown(rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {
	res := Run_python_drown(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	if res == true {
		Judge_drown(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
}

func Judge_fall(rtsp_data string, vid_stride int, threshold float32, where_loc string, get_camera *gin.Context) {
	res := Run_python_fall(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	if res == true {
		Judge_fall(rtsp_data, vid_stride, threshold, where_loc, get_camera)
	}
}

//以上代码中，timeout 是一个 time.After() 函数返回的 chan time.Time 类型的 channel，当超过设定时间后会从这个 channel 中接收到一个时间值，表示已经超时了。在 for 循环中，我们通过 select 语句来监听 timeout 和 ticker 这两个 channel，如果 timeout 先接收到了值，那么就退出循环，停止调用 Run_python_wave() 函数。如果 ticker 先接收到了值，那么就调用 Run_python_wave() 函数，然后等待下一次循环。这样就可以实现在超过一定时间后停止函数的调用。
