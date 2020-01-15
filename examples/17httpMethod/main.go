package main

import "github.com/gin-gonic/gin"

func main() {
	gin.DisableConsoleColor()

	router := gin.Default()
	httpGroup := router.Group("/http")

	httpGroup.GET("/someGet", func(c *gin.Context) {
		c.String(200, "getting")
	})
	httpGroup.POST("/somePost", func(c *gin.Context) {
		c.String(200, "posting")
	})
	httpGroup.PUT("/somePut", func(c *gin.Context) {
		c.String(200, "Putting")
	})
	httpGroup.DELETE("/someDelete", func(c *gin.Context) {
		c.String(200, "deleting")
	})
	httpGroup.PATCH("/somePatch", func(c *gin.Context) {
		c.String(200, "patching")
	})
	httpGroup.HEAD("/someHead", func(c *gin.Context) {
		c.String(200, "head")
	})
	httpGroup.OPTIONS("/someOptions", func(c *gin.Context) {
		c.String(200, "options")
	})

	router.Run(":8080")
}
