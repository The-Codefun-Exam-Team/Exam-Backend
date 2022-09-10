package models

import (
	"database/sql"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

type Problem struct {
	Pid int
	Sid int
	Code string
	Name string
	Type string
	Scoretype string
	Cid sql.NullInt32
	Status sql.NullString
	Pgroup string
	Statement string
	Timelimit float32
	Score float64
	Usechecker int
	Checkercode string
	Solved int
	Total int
}

type JSONProblem struct {
	Code string `json:"code"`
	Pid int `json:"id"`
	Name string `json:"name"`
}

func ReadProblemWithID(db *db.DB, id int) (Problem, error) {
	var prob Problem

	row := db.QueryRow("SELECT * FROM problems WHERE pid = ?", id)

	if err := row.Scan(&prob.Pid, &prob.Sid, &prob.Code, &prob.Name, &prob.Type, &prob.Scoretype, &prob.Cid, &prob.Status, &prob.Pgroup,
	&prob.Statement, &prob.Timelimit, &prob.Score, &prob.Usechecker, &prob.Checkercode, &prob.Solved, &prob.Total); err != nil {
		return prob, err
	}

	return prob, nil
}

func ReadJSONProblemWithID(db *db.DB, id int) (JSONProblem, error) {
	prob, err := ReadProblemWithID(db, id)
	if err != nil {
		var jprob JSONProblem
		return jprob, err
	}

	return JSONProblem{
		Code: prob.Code,
		Pid: prob.Pid,
		Name: prob.Name,
	}, err
}