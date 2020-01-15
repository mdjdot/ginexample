package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "Guest")
		lastName := c.Query("lastname")

		c.String(200, "Hello, %s %s", firstName, lastName)
	})
	router.Run(":8080")
}

// curl -v http://localhost:8080/welcome?lastname=kkzz&firstname=j%20d
// curl -v http://localhost:8080/welcome?lastname=kkzz
