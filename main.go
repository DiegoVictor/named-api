package main

import (
	"net/http"
	"os"

	"github.com/diegovictor/named-api/controllers"
	"github.com/diegovictor/named-api/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", []byte("Refer to the documentation: <a href=\"https://github.com/DiegoVictor/named-api\">here</a>."))
	})

	router.Use(middlewares.Cors())
	router.GET("/names", controllers.GetNames)
	router.POST("/datasets/:id/feedbacks", controllers.PostFeedback)

	router.StaticFile("/datasets", "./config/datasets.json")

	router.MaxMultipartMemory = 2 << 20
	router.POST("/upload", controllers.Upload)

	router.Run(":" + port)
}
