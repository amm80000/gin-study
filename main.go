package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/old", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com")
	})

	router.POST("/submit", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/result")
	})

	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/final"
		router.HandleContext(c)
	})

	router.GET("/final", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	router.GET("/result", func(c *gin.Context) {
		c.String(http.StatusOK, "Redirected here!")
	})

	err := router.Run()
	if err != nil {
		return
	}
}
