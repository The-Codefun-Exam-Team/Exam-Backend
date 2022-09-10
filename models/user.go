package models

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

type JSONVerifyUser struct {
	Tid int `json:"id"`
}

type JSONVerify struct {
	Data JSONVerifyUser `json:"data"`
	Valid bool
}

func Verify(db *db.DB, bearer_token string) (*JSONVerify, error) {
	var u JSONVerify
	u.Valid = false

	req, err := http.NewRequest(http.MethodGet, "https://codefun.vn/api/verify", nil)
	if err != nil {
		return &u, err
	}

	req.Header.Add("Authorization", bearer_token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Chrome/105.0.0.0")

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