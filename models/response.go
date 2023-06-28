package models

// Response is a struct that defines the format of the JSON response
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}
