package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
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

	routine.GET("/camera", detect_result.Search_camera)
	// get请求的同时发一个raw的json{
	//    "number" : "2021416365"
	//}

	routine.GET("/Result_to_front", detect_result.Result_to_front)
	//routine.GET("/result", detect_result.Get_camera)

	routine.GET("/wechat", detect_result.To_weixin_test)

	routine.GET("/to_flv1", detect_result.To_flv1)
	routine.GET("/to_flv2", detect_result.To_flv2)
	routine.GET("/to_flv3", detect_result.To_flv3)
	routine.GET("/to_flv4", detect_result.To_flv4)
	routine.GET("/to_flv5", detect_result.To_flv5)
	routine.GET("/to_flv6", detect_result.To_flv6)

	//GET请求，并发送一个raw的json
	//{
	//	rtsp_location="rtsp流地址"
	//}

	routine.Run("0.0.0.0:9000")
}
