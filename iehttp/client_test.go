package iehttp

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	client := NewClient()
	resp, err := client.Get("https://postman-echo.com/get?test=123", nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestAddHeader(t *testing.T) {
	testRequest, err := http.NewRequest(http.MethodGet, "https://postman-echo.com", nil)
	assert.NoError(t, err)
	assert.NotNil(t, testRequest)

	insertHeader := http.Header{}
	insertHeader["Auth"] = []string{"test", "test2"}
	testHeader := http.Header{}
	testHeader["Auth"] = []string{"test", "test2"}
	addHeader(testRequest, insertHeader)

	assert.EqualValues(t, testRequest.Header, testHeader)

}
