package main

import (
	"fmt"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("templates/raw.tmpl")

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(200, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2020, 1, 13, 15, 7, 0, 0, time.Local),
		})
	}) // Date:2020/01/13
	router.Run(":8080")

}
