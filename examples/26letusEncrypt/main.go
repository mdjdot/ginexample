package main

import "github.com/gin-gonic/gin"

import "log"
import "github.com/gin-gonic/autotls"
import "golang.org/x/crypto/acme/autocert"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// log.Fatal(autotls.Run(r, "example1.com", "example2.com"))

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example3.com"),
		Cache:      autocert.DirCache("/www/.cache"),
	}
	log.Fatal(autotls.RunWithManager(r, &m))
}
