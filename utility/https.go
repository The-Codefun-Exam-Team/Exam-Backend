package utility

import (
	"errors"
	"io"
	"net/http"
)

// ConstructRequest creates a request and add certain headers.
func ConstructRequest(method string, url string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Debug_Codefun/2.0")

	return request, nil
}

// ProcessRequest uses a http client to process the request, and return the response as []byte.
func ProcessRequest(request *http.Request) ([]byte, error) {
	rawResponse, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer rawResponse.Body.Close()

	if rawResponse.StatusCode != 200 {
		return nil, errors.New("non-200 status code")
	}

	response, err := io.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}
