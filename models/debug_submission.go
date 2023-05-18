package models

import (
	"github.com/jmoiron/sqlx"
)

// ShortenedSubmission contains brief information about a Debug submission.
type ShortenedSubmission struct {
	Drid       int     `json:"-" db:"drid"`
	Dpid       int     `json:"-" db:"dpid"`
	Dpcode     string  `json:"dpcode" db:"dpcode"`
	Rid        int     `json:"rid" db:"rid"`
	Tid        int     `json:"-" db:"tid"`
	Language   string  `json:"-" db:"language"`
	SubmitTime int     `json:"-" db:"submittime"`
	Result     string  `json:"result" db:"result"`
	Score      float64 `json:"score" db:"score"`
}

type DebugSubmission struct {
	ShortenedSubmission
	Difference int    `json:"-" db:"diff"`
	Code       string `json:"-" db:"code"`
}

func (submission *DebugSubmission) Write(db *sqlx.DB) (int, error) {
	result, err := db.NamedExec(`
		INSERT INTO debug_submissions
		(dpid, rid, tid, language, submittime, result, score, diff, code)
		VALUES (:dpid, :rid, :tid, :language, :submittime, :result, :score, :diff, :code)
	`, submission)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}
