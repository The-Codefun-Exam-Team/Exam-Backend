package models

import (
	"database/sql"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

type DebugSubmission struct {
	Drid       int
	Dpid       int
	Rid        int
	Tid        int
	Language   string
	Submittime int64
	Result     string
	Score      float64
	Diff       int
	Code       string
}

type JSONDebugSubmission struct {
	Dpcode   string  `json:"debug_problem_code"`
	Rid      int     `json:"codefun_id"`
	Score    float64 `json:"edit_result"`
	Diff     int     `json:"edit_score"`
	CFResult string  `json:"result"`
	CFScore  float64 `json:"score"`
}

func ReadDebugSubmission(db *db.DB, id int) (*DebugSubmission, error) {
	var sub DebugSubmission

	row := db.QueryRow("SELECT * FROM debug_submissions WHERE drid = ?", id)

	if err := row.Scan(&sub.Drid, &sub.Dpid, &sub.Rid, &sub.Tid, &sub.Language, &sub.Submittime, &sub.Result, &sub.Score, &sub.Diff, &sub.Code); err != nil {
		return nil, err
	}

	return &sub, nil
}

func WriteDebugSubmission(db *db.DB, sub *DebugSubmission) (int64, error) {
	res, err := db.Exec("INSERT INTO debug_submissions (dpid, rid, tid, language, submittime, result, score, diff, code) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		sub.Dpid, sub.Rid, sub.Tid, sub.Language, sub.Submittime, sub.Result, sub.Score, sub.Diff, sub.Code)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, nil
}

func UpdateDebugSubmission(db *db.DB, id int, result string, score float64) error {
	_, err := db.Exec("UPDATE debug_submissions SET result = ?, score = ? WHERE drid = ?", result, score, id)
	if err != nil {
		return err
	}
	return nil
}

func ReadJSONDebugSubmission(db *db.DB, id int) (*JSONDebugSubmission, error) {
	sub, err := ReadDebugSubmission(db, id)
	if err != nil {
		return nil, err
	}

	dprob, err := ReadDebugProblemWithID(db, sub.Dpid)
	if err != nil {
		return nil, err
	}

	run, err := ReadRun(db, sub.Rid)
	if err != nil {
		return nil, err
	}

	return &JSONDebugSubmission{
		Dpcode:   dprob.Code,
		Rid:      sub.Rid,
		Score:    sub.Score,
		Diff:     int(sub.Diff),
		CFResult: run.Result,
		CFScore:  run.Score,
	}, nil
}

func GetMaxScore(db *db.DB, dpid int, tid int) (float64, error) {
	var max_score sql.NullFloat64

	row := db.QueryRow("SELECT MAX(score) FROM debug_submissions WHERE dpid = ? AND tid = ?", dpid, tid)

	if err := row.Scan(&max_score); err != nil {
		return 0.0, err
	}

	if !max_score.Valid {
		return 0.0, nil
	}

	return max_score.Float64, nil
}
