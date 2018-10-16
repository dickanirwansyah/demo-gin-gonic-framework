package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context){
		c.String(http.StatusOK, "Hello Gin Framework")
	})
	r.GET("/ping", func(c *gin.Context){
		c.String(http.StatusOK, "PONG")
	})

	http.Handle("/", r)
	r.Run(":3000")
}
