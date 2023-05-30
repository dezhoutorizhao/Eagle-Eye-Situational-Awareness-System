package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strongwill.com/db"
	"strongwill.com/detect_result"
	"strongwill.com/detection"
	"strongwill.com/utils"
)

func main() {
	routine := gin.Default()
	routine.Use(db.Session("strongwill"))

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

	routine.POST("/user/login", db.Login, db.If_success)

	routine.POST("/user/register", db.AddUser_front, func(c *gin.Context) {
		db.Captcha(c, 4)
	})

	routine.GET("/captcha/verify/:value", func(c *gin.Context) {
		value := c.Param("value")
		if db.CaptchaVerify(c, value) {
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "failed"})
		}
	})

	routine.POST("/user/review", db.Review_func, db.Get_Review)

	routine.POST("/user/show", db.Show)
	// 传入一个{"school_id":"xxxid"}

	routine.POST("/camera/add_camera", detection.Add_camera)

	routine.POST("/camera/modify_camera", detection.Modify_camera)
	// 放在detection_db_init.go中了

	routine.GET("/camera", detect_result.Search_camera)
	// get请求的同时发一个raw的json{
	//    "number" : "2021416365"
	//}

	routine.GET("/Result_to_front", detect_result.Result_to_front)
	//routine.GET("/result", detect_result.Get_camera)

	//routine.GET("/wechat", detect_result.To_weixin_test)

	routine.POST("/Total_to_flv", detect_result.Total_to_flv)
	routine.GET("/Total_to_flv0", detect_result.To_flv0)
	routine.GET("/Total_to_flv1", detect_result.To_flv1)
	routine.GET("/Total_to_flv2", detect_result.To_flv2)
	routine.GET("/Total_to_flv3", detect_result.To_flv3)
	routine.GET("/Total_to_flv4", detect_result.To_flv4)
	routine.GET("/Total_to_flv5", detect_result.To_flv5)

	routine.POST("/Modify_logs_review", detect_result.Modify_logs_review)

	//GET请求，并发送一个raw的json
	//{
	//	rtsp_location="rtsp流地址"
	//}

	routine.GET("/Detect_process", detect_result.Detect_process)
	// 每添加一条记录重新调用一次

	routine.POST("/Trans_to_wechat", detect_result.To_weixin)

	routine.GET("/camera/All_camera", detection.All_camera)
	routine.GET("/user/All_users", db.All_users)
	routine.POST("/user/Modify_user", db.Modify_user)
	routine.GET("/Push_to_front", detect_result.Push_to_front)
	routine.POST("/user/Delete_user", db.Delete_user)
	routine.POST("/camera/Delete_camera", detection.Delete_camera)
	routine.GET("/show_to_flv", detect_result.Show_to_front)

	dir, _ := os.Getwd()
	fmt.Println(dir, "这是当前路径")
	fmt.Println(filepath.Join(dir, "/../flv"))

	//linux版本
	//routine.Static("/detect", filepath.Join(dir, "/detect"))
	//fmt.Println(filepath.Join(dir, "/detect"))
	//fmt.Println()
	//windows版本
	routine.Static("/detect", filepath.Join(dir, "/detect"))

	routine.Static("/Flv", filepath.Join(dir, "/../flv"))

	detect_result.Main_Starter()
	detect_result.Count_Location()
	//返回路由
	routine.GET("/Stop_detect", detect_result.Stop_Detect)

	routine.GET("/mqtt/subscribe", detect_result.Subscribe)
	routine.GET("/mqtt/un_subscribe", detect_result.Un_Subscribe)

	routine.Run("0.0.0.0:9000")
}
