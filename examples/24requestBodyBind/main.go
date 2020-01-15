package main

import "github.com/gin-gonic/gin"

import "github.com/gin-gonic/gin/binding"

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	// c.ShouldBind 使用了 c.Request.Body ，不可重用
	// c.
	// if errA := c.ShouldBind(&objA); errA == nil {
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(200, `the body should be formA`)
	}
	// if errB := c.ShouldBind(&objB); errB == nil {
	if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(200, `the body should be formB`)
	}
}

func main() {
	r := gin.Default()
	r.POST("/", SomeHandler)

	r.Run(":8080")
}
