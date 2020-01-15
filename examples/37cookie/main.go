package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 20, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})
	r.Run(":8080")
}
