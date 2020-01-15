package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 简单路由组：v1
	v1 := r.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.String(200, "login")
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.String(200, "submit")
		})
		v1.POST("/read", func(c *gin.Context) {
			c.String(200, "read")
		})
	}

	// 简单路由组：v2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.String(200, "login")
		})
		v2.POST("/submit", func(c *gin.Context) {
			c.String(200, "submit")
		})
		v2.POST("/read", func(c *gin.Context) {
			c.String(200, "read")
		})
	}

	r.Run(":8080")
}
