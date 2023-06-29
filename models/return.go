package models

// ReturnVerify contains the format of the "/verify" API from Codefun.
type ReturnVerify struct {
	User  User   `json:"data"`
	Error string `json:"error"`
}

// ReturnSubmission contains the format of the "/submissions/{id}" API from Codefun.
type ReturnSubmission struct {
	Submission CodefunSubmission `json:"data"`
	Error      string            `json:"error"`
}

