package models

import (
	"database/sql"
)

// ShortenedProblem contains brief information about a Debug problem.
type ShortenedProblem struct {
	// Code is the unique code of each problem.
	Code string `json:"code" db:"dpcode"`
	// BestScore is the best score of a user for the problem.
	// NULL score from RawBestScore is treated as 0.
	BestScore float64 `json:"best_score"`
	// RawBestScore is the score retrieved from the DB.
	RawBestScore sql.NullFloat64 `json:"-" db:"best_score"`
}

// DebugProblem contains all information about a Debug problem.
type DebugProblem struct {
	// DebugProblem inherits information from ShortenedProblem
	ShortenedProblem
	// Codetext is the code to be debugged for the problem.
	Codetext string `json:"codetext" db:"codetext"`
	// Judge contains information about all of the testcases.
	Judge Judge `json:"judge" db:"error"`
	// Language is the programming language that the given submission is in.
	Language string `json:"language" db:"language"`
	// Result is the verdict for the given submission.
	Result string `json:"result" db:"result"`
	// CodefunProblem is the associated problem for the given submission.
	CodefunProblem CodefunProblem `json:"problem" db:""`
}
