package main

import (
	"github.com/gin-gonic/gin"
	"strongwill.com/db"
	"strongwill.com/utils"
)

func main() {
	routine := gin.Default()

	// 登录注册方面

	// 获取cpu占用率
	routine.GET("occu/cpu",utils.Cpu_occu)
	// 获取内存占用率
	routine.GET("occu/mem",utils.Mem_occu)

	routine.POST("/login",db.Login)

	routine.Run("0.0.0.0:9000")
}