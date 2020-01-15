package main

import "github.com/gin-gonic/gin"

import "log"

func main() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r.POST("/foo", func(c *gin.Context) {
		c.JSON(200, "foo")
	})
	r.GET("/bar", func(c *gin.Context) {
		c.JSON(200, "bar")
	})
	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	r.Run(":8080")
}
