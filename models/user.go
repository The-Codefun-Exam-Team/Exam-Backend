package models

type User struct {
	ID          int     `json:"id"`
	Username    string  `json:"username"`
	DisplayName string  `json:"name"`
	Group       Group   `json:"group"`
	Status      string  `json:"status"`
	Avatar      string  `json:"avatar"`
	Score       float64 `json:"score"`
	Solved      int     `json:"solved"`
	Ratio       float64 `json:"ratio"`
	Email       string  `json:"email"`
}
