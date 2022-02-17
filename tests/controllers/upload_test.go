package tests

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/diegovictor/named-api/controllers"
	"github.com/diegovictor/named-api/helpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpload(t *testing.T) {
	t.Run("it should be able to upload a dataset", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "test.csv")
		part.Write([]byte("name\nDiego\nJulia\n"))
		writer.Close()

		request := httptest.NewRequest(http.MethodPost, "/", body)
		request.Header.Add("Content-Type", writer.FormDataContentType())

		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)
		context.Request = request

		controllers.Upload(context)

		assert.Equal(t, recorder.Code, 200)

		output := regexp.MustCompile("(\n|\r|\\s+)").ReplaceAllString(recorder.Body.String(), "")
		assert.Regexp(t, `{"id":\s?\d+}`, output)

		match := regexp.MustCompile(`\d+`).FindStringSubmatch(output)
		os.Remove(helpers.GetRoot() + "/datasets/" + match[0] + ".csv")
	})
}
