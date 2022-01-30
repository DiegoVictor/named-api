package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Catch(err error, c *gin.Context) bool {
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Ops! Something goes wrong, try again"})

		return true
	}

	return false
}
