package main

import "github.com/gin-gonic/gin"

import "github.com/gin-gonic/gin/binding"

import "net/http"

type myForm struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	err := c.ShouldBindWith(&fakeForm, binding.Form)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"colors": fakeForm.Colors})
}

func main() {
	r := gin.Default()

	r.POST("/", formHandler)

	r.Run(":8080")
}
