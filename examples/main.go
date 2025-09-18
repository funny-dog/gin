package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"message": "Gin Web Framework is running",
		})
	})

	// Hello World 端点
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Gin 信息端点
	r.GET("/gin/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"framework": "gin",
			"version": "v1.9.0",
			"status": "running",
		})
	})

	// 启动服务器，监听 8080 端口
	r.Run(":8080")
}