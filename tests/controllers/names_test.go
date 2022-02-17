package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"

	"github.com/diegovictor/named-api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetNames(t *testing.T) {
	t.Run("it should be able to return generated names", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		request := &http.Request{
			URL: &url.URL{},
		}

		query := request.URL.Query()
		query.Add("dataset", "1642602574998")
		request.URL.RawQuery = query.Encode()

		context.Request = request

		raw := []byte(`julia,pedro,diego,mariana,danilo,maria`)
		controllers.Command = func(name string, arg ...string) ([]byte, error) {
			return raw, nil
		}

		controllers.GetNames(context)

		assert.Equal(t, 200, recorder.Code)

		names := regexp.MustCompile("(\n|\"|\\[|]|\\s)").ReplaceAllString(recorder.Body.String(), "")
		assert.Equal(t, string(raw), names)
	})
}
