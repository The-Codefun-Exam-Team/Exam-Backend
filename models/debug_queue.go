package models

import (
	"database/sql"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/general"
)

type Queue struct {
	Rid  int
	Drid int
}

func AddToQueue(db *db.DB, q *Queue) error {
	_, err := db.Exec("INSERT INTO debug_queue (rid, drid) VALUES (?, ?)", q.Rid, q.Drid)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFromQueue(db *db.DB, id int) error {
	_, err := db.Exec("DELETE FROM debug_queue WHERE rid = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func ResolveQueue(db *db.DB) error {
	rows, err := db.Query("SELECT runs.rid, debug_queue.drid, runs.result, runs.score FROM runs INNER JOIN debug_queue ON runs.rid = debug_queue.rid WHERE runs.result NOT IN ('Q', 'R', '...')")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var q Queue
		var result string
		var score float64
		if err := rows.Scan(&q.Rid, &q.Drid, &result, &score); err != nil {
			return err
		}

		sub, err := ReadDebugSubmission(db, q.Drid)
		if err != nil {
			return err
		}

		dprob, err := ReadDebugProblemWithID(db, sub.Dpid)
		if err != nil {
			return err
		}

		final_score := general.CalculateScore(sub.Diff, score, dprob.Score, dprob.MinDiff)

		err = UpdateDebugSubmission(db, q.Drid, result, final_score)
		if err != nil {
			return err
		}

		err = DeleteFromQueue(db, q.Rid)
		if err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func Resolve1(db *db.DB, drid int) error {
	row := db.QueryRow("SELECT * FROM debug_queue WHERE drid = ?", drid)

	var q Queue
	if err := row.Scan(&q.Rid, &q.Drid); err != nil {
		if err == sql.ErrNoRows {
			return nil
		} else {
			return err
		}
	}

	run, err := ReadRun(db, q.Rid)
	if err != nil {
		return err
	}

	if run.Result == `Q` || run.Result == `R` || run.Result == `...` {
		return nil
	}

	sub, err := ReadDebugSubmission(db, q.Drid)
	if err != nil {
		return err
	}

	dprob, err := ReadDebugProblemWithID(db, sub.Dpid)
	if err != nil {
		return err
	}

	final_score := general.CalculateScore(sub.Diff, run.Score, dprob.Score, dprob.MinDiff)

	err = UpdateDebugSubmission(db, q.Drid, run.Result, final_score)
	if err != nil {
		return err
	}

	err = DeleteFromQueue(db, q.Rid)
	if err != nil {
		return err
	}

	return nil
}
