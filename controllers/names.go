package controllers

import (
	"net/http"
	"os/exec"
	"regexp"
	"strings"

	"github.com/diegovictor/named-api/helpers"

	"github.com/gin-gonic/gin"
)

func GetNames(c *gin.Context) {
	dataset := c.Query("dataset")

	params := []string{helpers.GetRoot() + "/lib/main.py", dataset}

	if !helpers.Catch(err, c) {
		text := helpers.ToUtf8(out)
		names := regexp.MustCompile("\r|\n").ReplaceAllString(text, "")
		result := strings.Split(names, ",")

		c.IndentedJSON(http.StatusOK, result)
	}
}
