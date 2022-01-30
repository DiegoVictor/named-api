package controllers

import (
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/diegovictor/name-generator/helpers"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")

	if !helpers.Catch(err, c) {
		id := time.Now().Unix()
		filename := fmt.Sprint(id)
		filePath := path.Join("./datasets", filename+".csv")

		err := c.SaveUploadedFile(file, filePath)

		if !helpers.Catch(err, c) {
			c.JSON(http.StatusOK, gin.H{"id": id})
		}
	}
}
