package main

import (
	"github.com/gin-gonic/gin"
	"strongwill.com/utils"
)

func main() {
	routine := gin.Default()

	routine.GET("/cpu",utils.Cpu_occu)
	routine.GET("/mem",utils.Mem_occu)

	routine.Run("0.0.0.0:9000")
}