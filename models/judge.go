package models

import (
	"errors"
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
	var judge_string string

	row := db.QueryRow("SELECT error FROM subs_code WHERE rid = ?", rid)

	if err := row.Scan(&judge_string); err != nil {
		return nil, err
	}

	j, err := ConvertToJudge(judge_string)
	if err != nil {
		// CE
		return &Judge{
			Correct: 0,
			Total: 0,
			Tests: []Test{
				{
					Verdict: "CE",
					RunningTime: 0.000,
					Message: judge_string,
				},
			},
		}, nil
	}

	return j, nil
}

func ConvertToJudge(raw string) (*Judge, error) {
	var tests []Test
	
	score_and_verdict := strings.Split(raw, "////")
	if len(score_and_verdict) != 2 {
		return nil, errors.New(`cannot split score and verdict`)
	}
	score, verdict := score_and_verdict[0], score_and_verdict[1]

	correct_and_total := strings.Split(score, "/")
	if len(correct_and_total) != 2 {
		return nil, errors.New(`cannot split to correct and total`)
	}
	strcorrect, strtotal := correct_and_total[0], correct_and_total[1]

	correct, err := strconv.Atoi(strcorrect)
	if err != nil {
		return nil, err
	}

	total, err := strconv.Atoi(strtotal)
	if err != nil {
		return nil, err
	}

	raw_tests := strings.Split(verdict, "||")
	for _, test := range raw_tests {
		result_time_error := strings.Split(test, "|")
		if len(result_time_error) != 3 {
			return nil, errors.New("not a valid testcase")
		}
		result, strtime, e := result_time_error[0], result_time_error[1], result_time_error[2]
		time, err := strconv.ParseFloat(strtime, 64)
		if err != nil {
			return nil, err
		}

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