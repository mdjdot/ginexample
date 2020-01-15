package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "<br>",
		}
		c.JSON(200, data) // //{"lang":"Go语言","tag":"\u003cbr\u003e"}
		// c.AsciiJSON(200, data) // {"lang":"Go\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data) // {"lang":"Go\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.PureJSON(http.StatusOK, data)  // {"lang":"Go语言","tag":"<br>"}
	})
	r.Run(":8080")
}
