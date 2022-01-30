package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/diegovictor/named-api/helpers"

	"github.com/gin-gonic/gin"
)

type Dataset struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
	Color    string `json:"color"`
	Title    string `json:"title"`
}

func GetDatasets(c *gin.Context) {
	data, err := ioutil.ReadFile("./config/datasets.json")

	if !helpers.Catch(err, c) {
		var parsedJson []Dataset

		json.Unmarshal([]byte(data), &parsedJson)
		c.IndentedJSON(http.StatusOK, parsedJson)
	}
}
