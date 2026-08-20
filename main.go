package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"method": "GET",
	})
}

func posting(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"method": "POST",
	})
}

func putting(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"method": "PUT",
	})
}

func deleting(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"method": "DELETE",
	})
}

func patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "PATCH",
	})
}

func head(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"method": "HEAD",
	})
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "OPTIONS",
	})
}

func main() {

	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatching", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)
	err := router.Run()
	if err != nil {
		return
	}
}
