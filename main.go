package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	//static
	r.Static("/assets", "./assets/")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "hjkim.dev",
		})
	})

	r.Run()
}
