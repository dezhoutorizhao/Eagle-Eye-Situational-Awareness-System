package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"strongwill.com/db"
	"strongwill.com/detection"
	"strongwill.com/utils"
)

func main() {
	routine := gin.Default()

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

	routine.GET("/camera/:number", detection.Serach_camera)

	routine.Run("0.0.0.0:9000")
}
