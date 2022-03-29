package controllers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/diegovictor/named-api/helpers"
	"github.com/diegovictor/named-api/services"

	"github.com/gin-gonic/gin"
)

var Command = services.Cli{}.Run

func GetNames(c *gin.Context) {
	dataset := c.Query("dataset")

	params := []string{"./lib/main.py", dataset}
	out, err := Command("python", params...)

	if !helpers.Catch(err, c) {
		text := helpers.ToUtf8(out)
		names := regexp.MustCompile("\r|\n").ReplaceAllString(text, "")
		result := strings.Split(names, ",")

		c.IndentedJSON(http.StatusOK, result)
	}
}
