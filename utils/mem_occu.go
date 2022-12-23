package utils

import (
	"github.com/gin-gonic/gin"
	"runtime"
)

// MemUsage represents the memory usage of the current process.
type MemUsage struct {
	Alloc uint64 `json:"alloc"` // 此字段表示程序分配的总字节数
	Total uint64 `json:"total"` // 自程序启动以来分配的总字节数
	NumAlloc uint64 `json:"num_alloc"` //  程序分配的内存总数
}

func Mem_occu(c *gin.Context) {
	// 获取当前程序的内存使用情况
	var usage MemUsage

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	//ReadMemStats函数的返回值，该值是一个近似值，旨在提供程序内存使用的粗略估计。

	usage.Alloc = memStats.Alloc
	usage.Total = memStats.TotalAlloc
	usage.NumAlloc = memStats.Mallocs

	// 将内存使用情况转换为 JSON 格式
	jsonData := MemUsage{usage.Alloc,usage.Total,usage.NumAlloc}

	c.JSON(200,jsonData)
}
