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
	Score    float64
	Result   string
}

type JSONDebugProblem struct {
	Problem  *JSONProblem `json:"problem"`
	Language string       `json:"language"`
	Result   string       `json:"result"`
	MaxScore float64      `json:"best_score"`
	Code     string       `json:"code"`
	Judge    *Judge       `json:"judge"`
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

func ReadJSONDebugProblemWithCode(db *db.DB, code string, tid int) (*JSONDebugProblem, error) {
	prob, err := ReadDebugProblemWithCode(db, code)
	if err != nil {
		return nil, err
	}

	jprob, err := ReadJSONProblemWithID(db, prob.Pid)
	if err != nil {
		return nil, err
	}

	codetext, err := ReadSubsCode(db, prob.Rid)
	if err != nil {
		return nil, err
	}

	judge, err := ReadJudge(db, prob.Rid)
	if err != nil {
		return nil, err
	}

	max_score, err := GetMaxScore(db, prob.Dpid, tid)
	if err != nil {
		return nil, err
	}

	return &JSONDebugProblem{
		Problem:  jprob,
		Language: prob.Language,
		Result:   prob.Result,
		MaxScore: max_score,
		Code:     codetext,
		Judge:    judge,
	}, nil
}
