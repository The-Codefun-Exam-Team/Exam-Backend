package models

import (
	"log"

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
	log.Print("Querying")

	rows, err := db.Query("SELECT runs.rid, debug_queue.drid FROM runs INNER JOIN debug_queue ON runs.rid = debug_queue.rid WHERE runs.result NOT IN ('Q', 'R', '...')")
	if err != nil {
		return err
	}

	log.Print("Got results")

	defer rows.Close()

	for rows.Next() {
		log.Print("Scanning")

		var q Queue
		if err := rows.Scan(&q.Rid, &q.Drid); err != nil {
			return err
		}

		log.Printf("Scanned rid = %v, drid = %v", q.Rid, q.Drid)

		log.Print("Reading run")

		run, err := ReadRun(db, q.Rid)
		if err != nil {
			return err
		}

		log.Print("Reading debug sub")

		sub, err := ReadDebugSubmission(db, q.Drid)
		if err != nil {
			return err
		}

		log.Print("Reading subs code")

		org_code, err := ReadSubsCode(db, sub.Rid)
		if err != nil {
			return err
		}

		log.Print("Formatting")

		org_len := len(general.Format(org_code))

		percentage := float64(sub.Diff) / float64(org_len)

		log.Printf("Diff: %v, Percentage: %v", sub.Diff, percentage)

		var final_score float64

		if run.Result != "AC" {
			final_score = 0
		} else {
			if percentage >= 80 {
				final_score = 100
			} else if percentage < 40 {
				final_score = 0
			} else {
				final_score = (percentage - 40) / (80 - 40)
			}
		}

		log.Printf("Final Score: %v", final_score)

		log.Print("Update debug sub")

		err = UpdateDebugSubmission(db, q.Drid, run.Result, final_score)
		if err != nil {
			return err
		}

		log.Print("Delete from queue")

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
