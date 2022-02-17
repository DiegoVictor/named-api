package tests

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/diegovictor/named-api/controllers"
	"github.com/diegovictor/named-api/helpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostFeedback(t *testing.T) {
	t.Run("it should be able to save the feedback", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		dataset := fmt.Sprint(helpers.GetRoot(), "/datasets/test.csv")

		os.Remove(dataset)
		file, _ := os.OpenFile(dataset, os.O_CREATE|os.O_RDWR, 0600)
		file.WriteString("name")
		file.Close()

		context.Request = httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader([]byte(`[{"name":"Diego","value":1},{"name":"Julia","value":-1}]`)),
		)
		context.Request.Header.Add("Content-Type", "application/json")
		context.Params = []gin.Param{
			{
				Key:   "id",
				Value: "test",
			},
		}

		controllers.PostFeedback(context)

		assert.Equal(t, "OK", recorder.Body.String())

		content, _ := ioutil.ReadFile(dataset)
		assert.Equal(t, "name\nDiego", string(content))
		os.Remove(dataset)
	})

	t.Run("it should not be able to retrieve dataset id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		context.Request = httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader([]byte(`[{"name":"Diego","value":1},{"name":"Julia","value":-1}]`)),
		)
		context.Request.Header.Add("Content-Type", "application/json")

		controllers.PostFeedback(context)
		assert.Equal(t, "{\n    \"message\": \"Ops! Something goes wrong, try again\"\n}", recorder.Body.String())
	})

	t.Run("it should not be able to open dataset file", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		context.Request = httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader([]byte(`[{"name":"Diego","value":1},{"name":"Julia","value":-1}]`)),
		)
		context.Request.Header.Add("Content-Type", "application/json")
		context.Params = []gin.Param{
			{
				Key:   "id",
				Value: "test",
			},
		}

		controllers.PostFeedback(context)
		assert.Equal(t, "{\n    \"message\": \"Ops! Something goes wrong, try again\"\n}", recorder.Body.String())
	})

	t.Run("it should not be able to write on dataset", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		dataset := fmt.Sprint(helpers.GetRoot(), "/datasets/fail.csv")

		os.Remove(dataset)
		file, _ := os.OpenFile(dataset, os.O_CREATE, 0600)
		file.Close()

		context.Request = httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader([]byte(`[{"name":"Diego","value":1},{"name":"Julia","value":-1}]`)),
		)
		context.Request.Header.Add("Content-Type", "application/json")
		context.Params = []gin.Param{
			{
				Key:   "id",
				Value: "fail",
			},
		}

		controllers.Append = func(file *os.File, text string) (int, error) {
			return 1, errors.New("Test error")
		}

		controllers.PostFeedback(context)
		assert.Equal(t, "{\n    \"message\": \"Ops! Something goes wrong, try again\"\n}", recorder.Body.String())
		os.Remove(dataset)
	})
}
