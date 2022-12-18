package utility

import (
	"errors"
	"io"
	"net/http"
)

func ConstructRequest(method string, url string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")

	return request, nil
}

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
