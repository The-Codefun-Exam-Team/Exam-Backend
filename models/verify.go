package models

// Verify contains the format of the "/verify" API from Codefun.
type Verify struct {
	User  User   `json:"data"`
	Error string `json:"error"`
}
