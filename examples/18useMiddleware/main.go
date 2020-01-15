package main

import "github.com/gin-gonic/gin"

var accounts = gin.Accounts{}

func main() {
	r := gin.Default()

	// Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.
	// By default gin.DefaultWriter = os.Stdout.
	r.Use(gin.Logger())

	// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/benchmark", gin.Logger(), gin.Recovery())

	// 认证路由组
	authorized := r.Group("/")
	authorized.Use(gin.BasicAuth(accounts))
	{
		authorized.POST("/login", func(c *gin.Context) {

		})
		authorized.POST("/submit", func(c *gin.Context) {

		})

		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analystics", func(c *gin.Context) {})
	}
	r.Run(":8080")
}
