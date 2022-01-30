package controllers

import (
	"net/http"
	"os"
	"strings"

	"github.com/diegovictor/named-api/helpers"

	"github.com/gin-gonic/gin"
)

type Names struct {
	Name  string
	Value int
}

type Params struct {
	Id string `uri:"id" binding:"required"`
}

func PostFeedback(c *gin.Context) {
	var names []Names

	if c.ShouldBind(&names) == nil {
		var params Params

		err := c.ShouldBindUri(&params)
		if helpers.Catch(err, c) {
			return
		}

		file, err := os.OpenFile("datasets/"+params.Id+".csv", os.O_APPEND|os.O_WRONLY, 0600)
		if helpers.Catch(err, c) {
			return
		}
		defer file.Close()

		for _, item := range names {
			if item.Value == 1 {
				_, err := file.WriteString("\n" + strings.Title(item.Name))
				if helpers.Catch(err, c) {
					return
				}
			}
		}
	}

	c.String(http.StatusOK, "OK")
}
