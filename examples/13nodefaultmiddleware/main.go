package main

import "github.com/gin-gonic/gin"

func main() {
	// // Default returns an Engine instance with the Logger and Recovery middleware already attached.
	// func Default() *Engine {
	// debugPrintWARNINGDefault()
	// engine := New()
	// engine.Use(Logger(), Recovery())
	// return engine
	// router := gin.Default()
	// }
	// router := gin.Default()
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hi, middleware")
	})

	router.Run(":8080")
}
