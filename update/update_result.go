package update

import (
	"encoding/json"
	"errors"
	"fmt"

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

// UpdateResult update a submission based on its ID
func (m *Module) UpdateResult(id string) error {
	// Fetch the status of that problem from the DB
	m.env.Log.Debugf("Trying to update submission %v", id)

	var status string
	var rid int
	row := m.env.DB.QueryRowx(getStatusQuery, id)
	err := row.Scan(&status, &rid)
	if err != nil {
		return err
	}

	if status != "Q" {
		// Submission does not need to be updated
		return nil
	}

	m.env.Log.Infof("Updating submission %v", id)

	// Fetch the submission from codefun
	m.env.Log.Infof("Fetching results for submission %v", rid)
	url := fmt.Sprintf("https://codefun.vn/api/submissions/%d", rid)

	m.env.Log.Debugf("Requesting from Codefun")
	request, err := utility.ConstructRequest("GET", url)
	if err != nil {
		return err
	}

	response, err := utility.ProcessRequest(m.env.Client, request)
	if err != nil {
		return err
	}

	m.env.Log.Debug("Unmarshalling response")
	var ret models.ReturnSubmission
	err = json.Unmarshal(response, &ret)
	if err != nil {
		return err
	}

	if ret.Error != "" {
		if ret.Error == fmt.Sprintf("Submission %d not found") {
			m.env.Log.Info("Submission not found")
			return nil
		}
		m.env.Log.Warnf("Error while fetching submission %v", id)
		return errors.New("Cannot fetch submission")
	}

	submission := ret.Submission

	// If the submission has not been evaluated, return
	notEvaluated := []string{"Q", "R", "..."}
	for _, status := range notEvaluated {
		if status == submission.Result {
			m.env.Log.Infof("Not yet evaluated")
			return nil
		}
	}

	// Update the submission
	return nil
}
