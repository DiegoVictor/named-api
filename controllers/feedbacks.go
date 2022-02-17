package controllers

import (
	"net/http"
	"strings"

	"github.com/diegovictor/named-api/helpers"
	"github.com/diegovictor/named-api/services"

	"github.com/gin-gonic/gin"
)

type Names struct {
	Name  string
	Value int
}

type Params struct {
	Id string `uri:"id" binding:"required"`
}

var (
	Append = services.File{}.Append
	Open   = services.File{}.Open
)

func PostFeedback(c *gin.Context) {
	var names []Names

	err := c.ShouldBind(&names)
	if !helpers.Catch(err, c) {
		var params Params

		err := c.ShouldBindUri(&params)
		if helpers.Catch(err, c) {
			return
		}

		file, err := Open(helpers.GetRoot() + "/datasets/" + params.Id + ".csv")
		if helpers.Catch(err, c) {
			return
		}
		defer file.Close()

		for _, item := range names {
			if item.Value == 1 {
				_, err := Append(file, strings.Title(item.Name))
				if helpers.Catch(err, c) {
					return
				}
			}
		}
	}

	c.String(http.StatusOK, "OK")
}
