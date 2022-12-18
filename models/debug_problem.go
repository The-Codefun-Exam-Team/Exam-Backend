package models

import (
	"database/sql"
)

type ShortenedProblem struct {
	Code         string          `json:"code" db:"dpcode"`
	BestScore    float64         `json:"best_score"`
	RawBestScore sql.NullFloat64 `json:"-" db:"best_score"`
}

type DebugProblem struct {
	ShortenedProblem
	Codetext       string         `json:"codetext" db:"codetext"`
	Judge          Judge          `json:"judge" db:"error"`
	Language       string         `json:"language" db:"language"`
	Result         string         `json:"result" db:"result"`
	CodefunProblem CodefunProblem `json:"problem" db:""`
}
