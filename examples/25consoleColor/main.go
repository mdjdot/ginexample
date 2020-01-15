package main

import "github.com/gin-gonic/gin"

func main() {
	// 禁用日志颜色
	// gin.DisableConsoleColor()

	// 强制启用日志颜色
	// gin.ForceConsoleColor()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
