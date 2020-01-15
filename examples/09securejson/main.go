package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.SecureJsonPrefix("forSecureJSON")
	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		c.SecureJSON(200, names)
	})
	r.Run(":8080")
}
