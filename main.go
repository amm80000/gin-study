package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {

		firstname := c.DefaultQuery("first_name", "John")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	err := router.Run()
	if err != nil {
		return
	}
}
