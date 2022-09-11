package models

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

type DebugSubmission struct {
	Drid       int
	Dpid       int
	Rid        int
	Tid        int
	Language   string
	Submittime int64
	Score      float64
	Diff       int
	Code       string
}

type JSONDebugSubmission struct {
	Dpid  int     `json:"debug_problem_id"`
	Rid   int     `json:"codefun_id"`
	Score float64 `json:"edit_result"`
	Diff  int     `json:"edit_score"`
}

func ReadDebugSubmission(db *db.DB, id int) (*DebugSubmission, error) {
	var sub DebugSubmission

	row := db.QueryRow("SELECT * FROM debug_submissions WHERE drid = ?", id)

	if err := row.Scan(&sub.Drid, &sub.Dpid, &sub.Rid, &sub.Tid, &sub.Language, &sub.Submittime, &sub.Score, &sub.Diff, &sub.Code); err != nil {
		return nil, err
	}

	return &sub, nil
}

func WriteDebugSubmission(db *db.DB, sub *DebugSubmission) (int64, error) {
	res, err := db.Exec("INSERT INTO debug_submissions (dpid, rid, tid, language, submittime, score, diff, code) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		sub.Dpid, sub.Rid, sub.Tid, sub.Language, sub.Submittime, sub.Score, sub.Diff, sub.Code)
	if err != nil {
		return 0, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, nil
}

func ReadJSONDebugSubmission(db *db.DB, id int) (*JSONDebugSubmission, error) {
	sub, err := ReadDebugSubmission(db, id)
	if err != nil {
		return nil, err
	}

	return &JSONDebugSubmission{
		Dpid:  sub.Dpid,
		Rid:   sub.Rid,
		Score: sub.Score,
		Diff:  int(sub.Diff),
	}, nil
}
