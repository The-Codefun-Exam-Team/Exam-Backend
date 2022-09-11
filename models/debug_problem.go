package models

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

type DebugProblem struct {
	Dpid     int
	Code     string
	Name     string
	Status   string
	Solved   int
	Total    int
	Rid      int
	Pid      int
	Language string
	Score    float32
	Result   string
}

type JSONDebugProblem struct {
	Problem  *JSONProblem `json:"problem"`
	Language string       `json:"language"`
	Result   string       `json:"result"`
	Score    float32      `json:"score"`
	Code     string       `json:"code"`
}

func ReadDebugProblemWithID(db *db.DB, dpid int) (*DebugProblem, error) {
	var prob DebugProblem

	row := db.QueryRow("SELECT * FROM debug_problems WHERE dpid = ?", dpid)

	if err := row.Scan(&prob.Dpid, &prob.Code, &prob.Name, &prob.Status, &prob.Solved, &prob.Total,
		&prob.Rid, &prob.Pid, &prob.Language, &prob.Score, &prob.Result); err != nil {
		return &prob, err
	}

	return &prob, nil
}

func ReadDebugProblemWithCode(db *db.DB, code string) (*DebugProblem, error) {
	var prob DebugProblem

	row := db.QueryRow("SELECT * FROM debug_problems WHERE code = ?", code)

	if err := row.Scan(&prob.Dpid, &prob.Code, &prob.Name, &prob.Status, &prob.Solved, &prob.Total,
		&prob.Rid, &prob.Pid, &prob.Language, &prob.Score, &prob.Result); err != nil {
		return &prob, err
	}

	return &prob, nil
}

func WriteDebugProblem(db *db.DB, prob *DebugProblem) (int64, error) {
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

func ReadSubsCode(db *db.DB, rid int) (string, error) {
	var r int
	var code, er string
	row := db.QueryRow("SELECT * FROM subs_code WHERE rid = ?", rid)

	if err := row.Scan(&r, &code, &er); err != nil {
		return code, err
	}

	return code, nil
}

func ReadJSONDebugProblemWithCode(db *db.DB, code string) (*JSONDebugProblem, error) {
	prob, err := ReadDebugProblemWithCode(db, code)
	if err != nil {
		var jdprob JSONDebugProblem
		return &jdprob, err
	}

	jprob, err := ReadJSONProblemWithID(db, prob.Pid)
	if err != nil {
		var jdprob JSONDebugProblem
		return &jdprob, err
	}

	codetext, err := ReadSubsCode(db, prob.Rid)
	if err != nil {
		var jdprob JSONDebugProblem
		return &jdprob, err
	}

	return &JSONDebugProblem{
		Problem:  jprob,
		Language: prob.Language,
		Result:   prob.Result,
		Score:    prob.Score,
		Code:     codetext,
	}, nil
}
