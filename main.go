package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	e "github.com/saleem92/go-package-example"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		e.Run()
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
