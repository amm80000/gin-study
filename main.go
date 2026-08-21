package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20

	router.POST("/upload", func(c *gin.Context) {

		form, err := c.MultipartForm()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		files := form.File["files"]

		for _, file := range files {

			log.Println(file.Filename)

			dst := filepath.Join("./files/", filepath.Base(file.Filename))

			c.SaveUploadedFile(file, dst)

			c.String(http.StatusOK, fmt.Sprintf("'%s' files uploaded!", file.Filename))
		}
	})

	err := router.Run()
	if err != nil {
		return
	}
}
