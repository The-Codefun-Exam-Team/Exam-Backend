package models

type Verify struct {
	User  User   `json:"data"`
	Error string `json:"error"`
}
