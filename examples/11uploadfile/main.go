package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(file.Filename)
		c.SaveUploadedFile(file, "uploads/"+file.Filename)
		c.String(200, "'%s' uploaded!", file.Filename)
	})
	router.Run(":8080")
}
