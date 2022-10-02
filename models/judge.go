package models

import (
	"strconv"
	"strings"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

type Test struct {
	Verdict     string  `json:"verdict"`
	RunningTime float64 `json:"runningTime"`
	Message     string  `json:"message"`
}

type Judge struct {
	Correct int    `json:"correct"`
	Total   int    `json:"total"`
	Tests   []Test `json:"tests"`
}

func ReadJudge(db *db.DB, rid int) (*Judge, error) {
	var tests []Test
	var judge_string string

	row := db.QueryRow("SELECT error FROM subs_code WHERE rid = ?", rid)

	if err := row.Scan(&judge_string); err != nil {
		return nil, err
	}

	score_and_verdict := strings.Split(judge_string, "////")
	score, verdict := score_and_verdict[0], score_and_verdict[1]

	correct_and_total := strings.Split(score, "/")
	strcorrect, strtotal := correct_and_total[0], correct_and_total[1]

	correct, _ := strconv.Atoi(strcorrect)
	total, _ := strconv.Atoi(strtotal)

	raw_tests := strings.Split(verdict, "||")
	for _, test := range raw_tests {
		result_time_error := strings.Split(test, "|")
		result, strtime, e := result_time_error[0], result_time_error[1], result_time_error[2]
		time, _ := strconv.ParseFloat(strtime, 64)

		tests = append(tests, Test{
			Verdict:     result,
			RunningTime: time,
			Message:     e,
		})
	}

	return &Judge{
		Correct: correct,
		Total:   total,
		Tests:   tests,
	}, nil
}
