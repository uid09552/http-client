package iehttp

import (
	"net/http"
)

func (c *ieClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	addHeader(request, headers)
	addHeader(request, c.DefaultHeaders)

	//must be called after all header were added to request
	parseBody(request.Header, body)
	return client.Do(request)
}

func parseBody(headers http.Header, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	return NewParser(headers.Get("content-type")).Parse(body)
}

//addHeader will add custom headers to application
func addHeader(r *http.Request, headers http.Header) {
	for key, values := range headers {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}
}
