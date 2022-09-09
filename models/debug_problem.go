package models

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

type DebugProblem struct {
	Dpid int `json:"-"`
	Code string `json:"-"`
	Name string `json:"-"`
	Status string `json:"-"`
	Solved int `json:"-"`
	Total int `json:"-"`
	Rid int `json:"-"`
	Pid int `json:"pid"`
	Language string `json:"language"`
	Score float32 `json:"score"`
	Result string `json:"result"`

	Codetext string `json:"code"`
}

func ReadDebugProblemID(db *db.DB, dpid int) (DebugProblem, error) {
	var prob DebugProblem

	row := db.QueryRow("SELECT * FROM debug_problems WHERE dpid = ?", dpid)

	if err := row.Scan(&prob.Dpid, &prob.Code, &prob.Name, &prob.Status, &prob.Solved, &prob.Total,
	&prob.Rid, &prob.Pid, &prob.Language, &prob.Score, &prob.Result); err != nil {
		return prob, err
	}

	return prob, nil
}

func ReadDebugProblemCode(db *db.DB, code string) (DebugProblem, error) {
	var prob DebugProblem

	row := db.QueryRow("SELECT * FROM debug_problems WHERE code = ?", code)

	if err := row.Scan(&prob.Dpid, &prob.Code, &prob.Name, &prob.Status, &prob.Solved, &prob.Total,
	&prob.Rid, &prob.Pid, &prob.Language, &prob.Score, &prob.Result); err != nil {
		return prob, err
	}

	return prob, nil
}

func WriteDebugProblem(db *db.DB, prob DebugProblem) (int64, error) {
	res, err := db.Exec("INSERT INTO debug_problems (code, name, status, solved, total, rid, pid, language, score, result) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	prob.Code, prob.Name, prob.Status, prob.Total, prob.Rid, prob.Pid, prob.Language, prob.Score, prob.Result)

	if err != nil {
		return 0, err
	}

	row_count, err := res.RowsAffected()

	if err != nil {
		return row_count, err
	}

	return row_count, nil
}