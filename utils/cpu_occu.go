package utils

import (
	"github.com/gin-gonic/gin"
	"runtime"
	"time"
)

// CPUUsage represents the CPU usage of the current process.
type CPUUsage struct {
	Percent float64 `json:"percent"`
	NumCpu int `json:"num_cpu"`
}

func Cpu_occu(c *gin.Context) {
	// 获取当前 CPU 核心数
	numCPU := runtime.NumCPU()

	// 获取当前程序的 CPU 使用情况
	var usage CPUUsage
	for i := 0; i < 10; i++ {
		t1 := time.Now()
		time.Sleep(time.Second)
		t2 := time.Now()

		cpu := runtime.NumCPU()
		elapsed := t2.Sub(t1).Seconds()
		usage.Percent += (float64(cpu) / float64(numCPU)) * 10 * elapsed
	}
	usage.Percent /= 10

	// 设置 NumCpu 字段的值
	usage.NumCpu = numCPU

	// 将 CPU 使用情况转换为 JSON 格式
	jsonData := CPUUsage{usage.Percent,usage.NumCpu}


	c.JSON(200, jsonData)
}