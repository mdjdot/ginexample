package main

import "github.com/gin-gonic/gin"

import "net/http"

func main() {
	r := gin.Default()
	r.GET("/tobaidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	r.GET("/index", func(c *gin.Context) {
		c.String(200, "welcome to index page")
	})
	r.GET("/toindex", func(c *gin.Context) {
		// c.Request.URL.Path = "/index"
		// r.HandleContext(c)
		c.Redirect(301, "/index")
	})

	r.Run(":8080")
}
