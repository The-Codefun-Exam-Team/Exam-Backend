package models

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

type Run struct {
	Rid        int
	Pid        int
	Tid        int
	Language   string
	Time       float64
	Result     string
	Access     string
	Submittime int
	Scored     int
	Score      float32
}

func ReadRun(db *db.DB, rid int) (Run, error) {
	var run Run

	row := db.QueryRow("SELECT * FROM runs WHERE rid = ?", rid)
	if err := row.Scan(&run.Rid, &run.Pid, &run.Tid, &run.Language, &run.Time, &run.Result, &run.Access,
		&run.Submittime, &run.Scored, &run.Score); err != nil {
		return run, err
	}

	return run, nil
}
