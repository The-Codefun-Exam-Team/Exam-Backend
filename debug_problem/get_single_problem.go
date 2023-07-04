package debugproblem

import (
	"database/sql"
	"net/http"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
	"github.com/labstack/echo/v4"
)

// Query for retrieving a single problem from DB.
var getSingleProblemQuery = `
SELECT

MAX(debug_submissions.score) AS best_score,
debug_problems.code AS dpcode,
debug_problems.name AS dpname,
debug_problems.language,
debug_problems.result,
subs_code.code AS codetext,
subs_code.error,
problems.*

FROM debug_problems
INNER JOIN subs_code ON subs_code.rid = debug_problems.rid
INNER JOIN problems ON problems.pid = debug_problems.pid
INNER JOIN debug_submissions ON debug_submissions.dpid = debug_problems.dpid
	AND debug_submissions.tid = ?

WHERE debug_problems.code = ?
`

// GetSingleProblem returns a DebugProblem
func (m *Module) GetSingleProblem(c echo.Context) (err error) {
	// Verify the user first
	user, err := utility.Verify(c, m.env)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}

	// Getting the problem code
	code := c.Param("code")
	m.env.Log.Infof("Getting problem (%v)", code)

	var userID int

	if user == nil {
		userID = 0
	} else {
		userID = user.ID
	}

	// Query the DB
	m.env.Log.Debug("Querying DB for problem")
	var p models.DebugProblem
	err = m.env.DB.Get(&p, getSingleProblemQuery, userID, code)

	// Convert the score from NULL to 0
	m.env.Log.Debug("Converting score")
	if !p.RawBestScore.Valid {
		p.BestScore = 0.0
	} else {
		p.BestScore = p.RawBestScore.Float64
	}

	// Log errors
	if err != nil {
		if err == sql.ErrNoRows {
			// If no row was found
			m.env.Log.Infof("Getting problem: (%v) not found", code)
			return c.JSON(http.StatusNotFound, models.Response{
				Error: "No problem was found",
			})
		} else {
			// Another error
			m.env.Log.Errorf("Getting problem: Error encountered: %v", err)
			return c.JSON(http.StatusInternalServerError, models.Response{
				Error: "An error has occured",
			})
		}
	}

	// Return the problem
	m.env.Log.Infof("Found problem (%v)", code)
	return c.JSON(http.StatusOK, models.Response{
		Data: p,
	})
}
