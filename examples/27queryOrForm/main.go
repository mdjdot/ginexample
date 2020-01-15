package main

import "github.com/gin-gonic/gin"

import "fmt"

func main() {

	r := gin.Default()
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	r.Run(":8080")
}

// curl -v --form names[first]=thinkerou --form names[second]=tianou http://localhost:8080/post?ids[1]=123&ids[2]=456
