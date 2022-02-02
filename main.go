package main

import (
	"os"

	"github.com/diegovictor/named-api/controllers"
	"github.com/diegovictor/named-api/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	router := gin.Default()

	router.Use(middlewares.Cors())
	router.GET("/names", controllers.GetNames)
	router.POST("/datasets/:id/feedbacks", controllers.PostFeedback)

	router.StaticFile("/datasets", "./config/datasets.json")

	router.MaxMultipartMemory = 2 << 20
	router.POST("/upload", controllers.Upload)

	router.Run(":" + port)
}
