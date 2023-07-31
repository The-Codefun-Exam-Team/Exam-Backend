package models

import (
	"errors"
	"strconv"
	"strings"
)

// Test is a struct containing information about a testcase.
// It includes the verdict (AC, WA, ...), the execution time in seconds and the message for each testcase.
type Test struct {
	Verdict     string  `json:"verdict"`
	RunningTime float64 `json:"runningTime"`
	Message     string  `json:"message"`
}

// Judge is a struct containing information about multiple testcases of a submission.
// It includes the number of total and correct testcases, as well as a slice of all tests.
type Judge struct {
	CorrectTestCount int    `json:"correct"`
	TotalTestCount   int    `json:"total"`
	Tests            []Test `json:"tests"`
}

// Function for scanning Test
func (t *Test) Scan(src interface{}) (err error) {
	var source string

	// Converting data to string
	switch src.(type) {
	case string:
		source = src.(string)
	case []byte:
		source = string(src.([]byte))
	default:
		return errors.New("incompatible type for Test")
	}

	// Split the testcase into 3 parts:
	// Verdict, RunningTime and Message
	var temp []string

	temp = strings.Split(source, "|")
	if len(temp) != 3 {
		return errors.New("not a valid testcase")
	}

	var runningtime string

	t.Verdict, runningtime, t.Message = temp[0], temp[1], strings.TrimSpace(temp[2])

	// Convert RunningTime to float64
	t.RunningTime, err = strconv.ParseFloat(runningtime, 64)
	if err != nil {
		return
	}

	return nil
}

func (j *Judge) CompileError(message string) (err error) {
	j.CorrectTestCount = 0
	j.TotalTestCount = 0
	j.Tests = []Test{
		{
			Verdict: "CE",
			RunningTime: 0.000,
			Message: strings.TrimSpace(message),
		},
	}

	return nil
}

// Function for scanning Judge
func (j *Judge) Scan(src interface{}) (err error) {
	var source string

	// Convert data to string
	switch src.(type) {
	case string:
		source = src.(string)
	case []byte:
		source = string(src.([]byte))
	default:
		return errors.New("incompatible type for Judge")
	}

	// If the data starts with a path (with the character '/'), it means the verdict is CE.
	// Currently, when a submission is CE, this code will treat it as a Judge with 0/0 tests.
	// Tests will contain a single Test, with the verdict CE and the compiler message.
	if strings.HasPrefix(source, "/") {
		return j.CompileError(source)
	}

	// Split the judge string into the score (correct/total) and all of the testcases
	var temp []string

	temp = strings.Split(source, "////")
	if len(temp) != 2 {
		return j.CompileError(source)
	}

	score, verdict := temp[0], temp[1]

	// Split the score into correct count and total count
	temp = strings.Split(score, "/")
	if len(temp) != 2 {
		return j.CompileError(source)
	}

	correct, total := temp[0], temp[1]

	j.CorrectTestCount, err = strconv.Atoi(correct)
	if err != nil {
		return j.CompileError(source)
	}

	j.TotalTestCount, err = strconv.Atoi(total)
	if err != nil {
		return j.CompileError(source)
	}

	// Split all of the testcases, and use the Scan method of Test to process each one
	raw_tests := strings.Split(verdict, "||")

	for _, test := range raw_tests {
		var t Test
		t.Scan(test)
		j.Tests = append(j.Tests, t)
	}

	return nil
}
