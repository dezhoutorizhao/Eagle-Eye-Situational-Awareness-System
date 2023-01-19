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

	//// 设置允许跨域的域名列表
	//var allowOrigins = []string{"http://www.example.com", "http://www.example.org"}
	//
	//// 设置允许跨域的请求头
	//var allowHeaders = []string{"Content-Type", "X-Requested-With"}
	//
	//// 设置允许跨域的请求方法
	//var allowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	//
	//// 设置允许跨域的响应头
	//var exposeHeaders = []string{"Content-Length"}
	//
	//// 创建一个处理函数，用于处理跨域请求
	//func handler(w http.ResponseWriter, r *http.Request) {
	//	// 设置允许跨域的域名
	//	w.Header().Set("Access-Control-Allow-Origin", strings.Join(allowOrigins, ","))
	//	// 设置允许跨域的请求头
	//	w.Header().Set("Access-Control-Allow-Headers", strings.Join(allowHeaders, ","))
	//	// 设置允许跨域的请求方法
	//	w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowMethods, ","))
	//	// 设置允许跨域的响应头
	//	w.Header().Set("Access-Control-Expose-Headers", strings.Join(exposeHeaders, ","))
	//}
	//http.HandleFunc("/", handler)
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

	routine.POST("/add_camera", detection.Add_camera)

	routine.POST("/modify_camera", detection.Modify_camera)

	routine.GET("/camera", detect_result.Search_camera)

	routine.Run("0.0.0.0:9000")
}
