package main

import (
	"VOL/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 定义路由和处理函数
	router.POST("/k8s/command", handler.HandleK8sCommand)

	// 启动服务器
	router.Run(":10010")
}
