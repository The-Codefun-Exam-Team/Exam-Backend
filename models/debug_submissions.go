package models

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

type DebugSubmission struct {
	Drid       int
	Dpid       int
	Tid        int
	Language   string
	Submittime int64
	Score      float64
	Diff       float64
}

func ReadDebugSubmission(db *db.DB, id int) (*DebugSubmission, error) {
	var sub DebugSubmission

	row := db.QueryRow("SELECT * FROM debug_submissions WHERE drid = ?", id)

	if err := row.Scan(&sub.Drid, &sub.Dpid, &sub.Tid, &sub.Language, &sub.Submittime, &sub.Score, &sub.Diff); err != nil {
		return &sub, err
	}

	return &sub, nil
}

func WriteDebugSubmission(db *db.DB, sub *DebugSubmission) (int64, error) {
	res, err := db.Exec("INSERT INTO debug_submissions (dpid, tid, language, submittime, score, diff) VALUES (?, ?, ?, ?, ?, ?)",
	sub.Dpid, sub.Tid, sub.Language, sub.Submittime, sub.Score, sub.Diff)
	if err != nil {
		return 0, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, nil
}