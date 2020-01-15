package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 结构体
// 绑定JSON、XML、FORM
type Login struct {
	User     string `json:"user" form:"user" xml:"user" binding:"required"`
	Password string `json:"password" form:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// 绑定JSON
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "mamu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.Run(":8080")
}
