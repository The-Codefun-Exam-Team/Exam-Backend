package models

type CodefunSubmission struct {
	Rid         int            `json:"id"`
	Problem     CodefunProblem `json:"problem"`
	Owner       User           `json:"owner"`
	Language    string         `json:"language"`
	Result      string         `json:"result"`
	RunningTime float64        `json:"runningTime"`
	SubmitTime  int64          `json:"submitTime"`
	IsScored    bool           `json:"isScored"`
	Score       float64        `json:"score"`
	Code        string         `json:"code"`
	Tests       Judge          `json:"judge"`
}
