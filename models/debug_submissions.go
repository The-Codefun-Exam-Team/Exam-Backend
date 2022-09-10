package models

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

type DebugSubmission struct {
	Drid       int
	Dpid       int
	Tid        int
	Language   string
	Submittime int
	Score      float64
	Diff       float64
}

func ReadDebugSubmission(db *db.DB, id int) (DebugSubmission, error) {
	var sub DebugSubmission

	row := db.QueryRow("SELECT * FROM debug_submissions WHERE drid = ?", id)

	if err := row.Scan(&sub.Drid, &sub.Dpid, &sub.Tid, &sub.Language, &sub.Submittime, &sub.Score, &sub.Diff); err != nil {
		return sub, err
	}

	return sub, nil
}
