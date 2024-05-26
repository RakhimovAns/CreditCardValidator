package main

import (
	"github.com/RakhimovAns/CreditCardValidator/handlers"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {
	startServer()
}

func startServer() {
	r := gin.Default()
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.File(filepath.Join("static", "index.html"))
	})
	r.GET("/ping", handlers.Ping)
	r.POST("/check", handlers.CheckCard)
	r.Run()
}
