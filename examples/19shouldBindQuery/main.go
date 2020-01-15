package main

import "github.com/gin-gonic/gin"

import "log"

// Person .
type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("========Only Bind By Query String========")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}

func main() {
	router := gin.Default()
	router.Any("/testing", startPage)
	router.Run(":8080")
}
