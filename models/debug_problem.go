package models

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

// ShortenedProblem contains brief information about a Debug problem.
type ShortenedProblem struct {
	// Dpid is the debug problem's ID.
	Dpid int `json:"-" db:"dpid"`
	// Code is the unique code of each problem.
	Code string `json:"code" db:"dpcode"`
	// Name is the long name of each problem.
	Name string `json:"name" db:"dpname"`
	// Status is the status of the problem.
	// It can be "Active", "Deleted" or "Hidden"
	Status string `json:"-" db:"status"`
	// SolvedCount is the number of user that have got Accepted for each problem.
	SolvedCount int `json:"-" db:"solved"`
	// TotalCount is the number of attempt for each problem.
	TotalCount int `json:"-" db:"total"`
	// Language is the programming language that the given submission is in.
	Language string `json:"language" db:"language"`
	// Result is the verdict for the given submission.
	Result string `json:"result" db:"result"`
	// Rid is the ID of the submission associated with the problem.
	Rid int `json:"-" db:"rid"`
	// MinimumDifference is the smallest calculated distance between
	// the best solution and the given submission.
	MinimumDifference int `json:"-" db:"mindiff"`
	// BestScore is the best score of a user for the problem.
	// NULL score from RawBestScore is treated as 0.
	BestScore float64 `json:"best_score"`
	// RawBestScore is the score retrieved from the DB.
	RawBestScore sql.NullFloat64 `json:"-" db:"best_score"`
	// Pid is the ID of CodefunProblem associated with this
	Pid int `json:"-" db:"_pid"`
	// Score is the score of the submission associated with this
	Score float64 `json:"-" db:"_score"`
}

// DebugProblem contains all information about a Debug problem.
type DebugProblem struct {
	// DebugProblem inherits information from ShortenedProblem
	ShortenedProblem
	// Codetext is the code to be debugged for the problem.
	Codetext string `json:"codetext" db:"codetext"`
	// Judge contains information about all of the testcases.
	Judge Judge `json:"judge" db:"error"`
	// CodefunProblem is the associated problem for the given submission.
	CodefunProblem CodefunProblem `json:"problem" db:""`
}

func (dprob *DebugProblem) Write(db *sqlx.DB) (int, error) {
	result, err := db.NamedExec(`
		INSERT INTO debug_problems
		(code, name, solved, total, rid, pid, language, score, result, mindiff)
		VALUES (:dpcode, :dpname, :solved, :total, :rid, :_pid, :language, :_score, :result, :mindiff)
	`, dprob)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}
