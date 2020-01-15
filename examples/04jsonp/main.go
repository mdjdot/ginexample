package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/JSONP", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"foo": "bar",
		})
	})
	r.Run(":8080")
}

// http://localhost:8080/JSONP?callback=abc/cal
// abc/cal({"foo":"bar"});
