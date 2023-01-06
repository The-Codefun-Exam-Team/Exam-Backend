package debugsubmission

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
	"github.com/labstack/echo/v4"
)

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

	// TODO(unknown): Get the submission

	return
}
