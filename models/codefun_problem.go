package models

import (
	"database/sql"
)

// CodefunProblem is a struct containing all information about a problem from Codefun.
type CodefunProblem struct {
	// Pid is the problem's ID.
	Pid int `json:"pid" db:"pid"`
	// Sid is ID of something?
	Sid int `json:"-" db:"sid"`
	// ProblemCode is the unique code (eg. P001) of each problem.
	ProblemCode string `json:"code" db:"code"`
	// ProblemName is the full name of each problem.
	ProblemName string `json:"name" db:"name"`
	// ProblemType is an comma-separated string of the problem's tags.
	ProblemType string `json:"-" db:"type"`
	// ProblemScoreType is the scoring format of each problem ("oi" or "acm").
	ProblemScoreType string `json:"-" db:"scoretype"`
	// Cid is the ID of the contest each problem belongs to.
	Cid sql.NullInt32 `json:"-" db:"cid"`
	// Status is the status of each problem.
	// It is either "Active", "Deleted" or "Hidden"
	Status sql.NullString `json:"-" db:"status"`
	// ProblemGroup is the category of each problem.
	ProblemGroup string `json:"-" db:"pgroup"`
	// Statement is the statement for each problem.
	Statement string `json:"-" db:"statement"`
	// TimeLimit is the maximum execution time allowed for each problem, measured in seconds.
	TimeLimit float64 `json:"-" db:"timelimit"`
	// Score is the "maximum" score allowed for a submission for each problem.
	Score float64 `json:"-" db:"score"`
	// UseChecker is an integer with the value of 0 or 1.
	// It shows whether a checker is needed for each problem.
	UseChecker int `json:"-" db:"usechecker"`
	// CheckerCode is the code to be used for the checker, in case UseChecker = 1.
	CheckerCode string `json:"-" db:"checkercode"`
	// SolvedCount is the number of user that have got Accepted for each problem.
	SolvedCount int `json:"-" db:"solved"`
	// TotalCount is the number of attempt for each problem.
	TotalCount int `json:"-" db:"total"`
}
