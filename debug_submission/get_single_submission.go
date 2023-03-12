package debugsubmission

import (
	"database/sql"
	"net/http"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
	"github.com/labstack/echo/v4"
)

var getSingleSubmissionQuery = `
SELECT

debug_submissions.*,
debug_problems.code AS dpcode

FROM debug_submissions
INNER JOIN debug_problems ON debug_problems.dpid = debug_submissions.dpid

WHERE debug_submissions.drid = ?
`

// GetSingleSubmission returns a DebugSubmission
func (m *Module) GetSingleSubmission(c echo.Context) (err error) {
	// Verify the user first
	user, err := utility.Verify(c, m.env)
	if user == nil {
		return err
	}

	// Getting the submission ID
	id := c.Param("id")
	m.env.Log.Infof("Getting submission (%v)", id)

	// Query the DB
	m.env.Log.Debug("Querying DB for submission")
	var sub models.DebugSubmission
	err = m.env.DB.Get(&sub, getSingleSubmissionQuery, id)

	// Log errors
	if err != nil {
		if err == sql.ErrNoRows {
			// If no row was found
			m.env.Log.Infof("Getting submission: (%v) not found", id)
			return c.JSON(http.StatusNotFound, models.Response{
				Error: "No submission was found",
			})
		} else {
			// Another error
			m.env.Log.Errorf("Getting submission: Error encountered: %v", err)
			return c.JSON(http.StatusInternalServerError, models.Response{
				Error: "An error has occured",
			})
		}
	}

	// Return the submission
	m.env.Log.Infof("Found submission (%v)", id)
	return c.JSON(http.StatusOK, models.Response{
		Data: sub,
	})
}
