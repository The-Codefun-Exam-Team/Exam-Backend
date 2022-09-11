package models

import (
	"encoding/json"
	"io"
	"net/http"
)

type JSONVerifyUser struct {
	Tid int `json:"id"`
}

type JSONVerify struct {
	Data  *JSONVerifyUser `json:"data"`
	Valid bool
}

func Verify(bearer_token string) (*JSONVerify, error) {
	var u JSONVerify
	u.Valid = false

	req, err := http.NewRequest(http.MethodPost, "https://codefun.vn/api/verify", nil)
	if err != nil {
		return &u, err
	}

	req.Header.Add("Authorization", bearer_token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")

	rawresp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &u, err
	}

	defer rawresp.Body.Close()

	if rawresp.StatusCode != 200 {
		return &u, nil
	}

	body, err := io.ReadAll(rawresp.Body)
	if err != nil {
		return &u, err
	}

	var resp JSONVerify
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return &u, err
	}
	resp.Valid = true

	return &resp, nil
}
