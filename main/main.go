package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"runtime"
	"strongwill.com/db"
	"strongwill.com/detect_result"
	"strongwill.com/detection"
	"strongwill.com/utils"
)

func main() {
	routine := gin.Default()

	routine.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 假设响应数据为data，其长度为len(data)
		//c.Header("Content-Length", strconv.Itoa(100000000000000000))
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	//为解决wrote more than the declared Content-Length问题，使用流式数据的传输方式
	routine.GET("/stream_data", func(c *gin.Context) {
		// 假设响应数据为一个非常大的字符串
		data := "very large string ..."
		c.Stream(func(w io.Writer) bool {
			_, err := w.Write([]byte(data))
			if err != nil {
				return false
			}
			return true
		})
	})

	// 登录注册方面
	fmt.Println(runtime.Version())
	// 获取cpu占用率
	routine.GET("occu/cpu", utils.Cpu_occu)
	// 获取内存占用率
	routine.GET("occu/mem", utils.Mem_occu)

	routine.POST("/login", db.Login, db.If_success)

	routine.POST("/register", db.AddUser_front)

	routine.POST("/review", db.Review_func, db.Get_Review)

	routine.POST("/show", db.Show)
	// 传入一个{"school_id":"xxxid"}

	routine.POST("/add_camera", detection.Add_camera)

	routine.POST("/modify_camera", detection.Modify_camera)
	// 放在detection_db_init.go中了

	routine.GET("/camera", detect_result.Search_camera)
	// get请求的同时发一个raw的json{
	//    "number" : "2021416365"
	//}

	routine.GET("/Result_to_front", detect_result.Result_to_front)
	//routine.GET("/result", detect_result.Get_camera)

	//routine.GET("/wechat", detect_result.To_weixin_test)

	routine.GET("/Total_to_flv", detect_result.Total_to_flv)

	routine.POST("/Modify_logs_review", detect_result.Modify_logs_review)

	//GET请求，并发送一个raw的json
	//{
	//	rtsp_location="rtsp流地址"
	//}

	routine.GET("/Detect_process", detect_result.Detect_process)
	// 每添加一条记录重新调用一次

	routine.POST("/Trans_to_wechat", detect_result.To_weixin)

	routine.GET("/All_camera", detection.All_camera)
	routine.GET("/All_users", db.All_users)
	routine.POST("/Modify_user", db.Modify_user)
	routine.GET("/Push_to_front", detect_result.Push_to_front)
	routine.POST("/Delete_user", db.Delete_user)
	routine.POST("/Delete_camera", detection.Delete_camera)

	routine.Run("0.0.0.0:9000")
}
