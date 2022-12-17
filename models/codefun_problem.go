package models

import (
	"database/sql"
)

type CodefunProblem struct {
	Pid              int            `json:"pid" db:"pid"`
	Sid              int            `json:"-" db:"sid"`
	ProblemCode      string         `json:"code" db:"code"`
	ProblemName      string         `json:"name" db:"name"`
	ProblemType      string         `json:"-" db:"type"`
	ProblemScoreType string         `json:"-" db:"scoretype"`
	Cid              sql.NullInt32  `json:"-" db:"cid"`
	Status           sql.NullString `json:"-" db:"status"`
	ProblemGroup     string         `json:"-" db:"pgroup"`
	Statement        string         `json:"-" db:"statement"`
	TimeLimit        float64        `json:"-" db:"timelimit"`
	Score            float64        `json:"-" db:"score"`
	UseChecker       int            `json:"-" db:"usechecker"`
	CheckerCode      string         `json:"-" db:"checkercode"`
	SolvedCount      int            `json:"-" db:"solved"`
	TotalCount       int            `json:"-" db:"total"`
}
