package update

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/envlib"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
)

var getStatusQuery = `
SELECT

result,
rid

FROM debug_submissions
WHERE drid = ?
`

var getDPQuery = `
SELECT

subs_code.score,
subs_code.code AS codetext,
debug_problems.mindiff

FROM debug_submissions
INNER JOIN debug_problems ON debug_submissions.dpid = debug_problems.dpid
INNER JOIN subs_code ON debug_problems.rid = subs_code.rid

WHERE debug_submissions.drid = ?
`

var updateSubmissionExec = `
UPDATE

debug_submissions

SET

result = ?,
score = ?

WHERE

drid = ?
`

// UpdateResult update a submission based on its ID
func UpdateResult(env *envlib.Env, id string) error {
	// Fetch the status of that problem from the DB
	env.Log.Debugf("Trying to update submission %v", id)

	var status string
	var rid int
	row := env.DB.QueryRowx(getStatusQuery, id)
	err := row.Scan(&status, &rid)
	if err != nil {
		return err
	}

	if status != "Q" {
		// Submission does not need to be updated
		env.Log.Infof("Submission does not need to be updated")
		return nil
	}

	env.Log.Infof("Updating submission %v", id)

	// Fetch the submission from codefun
	env.Log.Infof("Fetching results for submission %v", rid)
	url := fmt.Sprintf("https://codefun.vn/api/submissions/%d", rid)

	env.Log.Debugf("Requesting from Codefun")
	request, err := utility.ConstructRequest("GET", url)
	if err != nil {
		return err
	}

	response, err := utility.ProcessRequest(env.Client, request)
	if err != nil {
		return err
	}

	env.Log.Debug("Unmarshalling response")
	var ret models.ReturnSubmission
	err = json.Unmarshal(response, &ret)
	if err != nil {
		return err
	}

	if ret.Error != "" {
		if ret.Error == fmt.Sprintf("Submission %d not found") {
			env.Log.Info("Submission not found")
			return nil
		}
		env.Log.Warnf("Error while fetching submission %v", id)
		return errors.New("Cannot fetch submission")
	}

	submission := ret.Submission

	// If the submission has not been evaluated, return
	notEvaluated := []string{"Q", "R", "..."}
	for _, status := range notEvaluated {
		if status == submission.Result {
			env.Log.Infof("Not yet evaluated")
			return nil
		}
	}

	// Update the submission

	var original_score float64
	var codetext string
	var mindiff int

	env.Log.Debug("Querying DB for information")
	row = env.DB.QueryRowx(getDPQuery, id)
	err = row.Scan(&original_score, &codetext, &mindiff)
	if err != nil {
		return err
	}

	env.Log.Debug("Evaluating submission")
	diff := utility.EditDistance(submission.Code, codetext)
	evaluation := utility.CalculateScore(diff, mindiff, submission.Score, original_score)

	// Write submission to DB
	// Update result and score based on drid

	var new_result string
	epsilon := 1e-6

	if evaluation+epsilon >= 100 {
		new_result = "AC"
	} else {
		new_result = "SS"
	}

	env.Log.Info("Writing submision %v to DB", id)
	_, err = env.DB.Exec(updateSubmissionExec, new_result, evaluation, id)
	if err != nil {
		return err
	}

	return nil
}
