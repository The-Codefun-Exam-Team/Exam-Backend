package submit

import (
	"encoding/json"
	"io"
	"net/http"
)

type SubmitResponse struct {
	Rid int `json:"data"`
}

type SubmitReturnValue struct {
	Drid int `json:"id"`
}

func ExtractResponse(response *http.Response) (int, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return -1, err
	}

	var resp SubmitResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return -1, err
	}

	return resp.Rid, nil
}
