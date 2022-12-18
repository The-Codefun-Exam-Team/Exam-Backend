package debugproblem

import (
	"database/sql"
	"net/http"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
	"github.com/labstack/echo/v4"
)

var getSingleProblemQuery = `
SELECT

MAX(debug_submissions.score) AS best_score,
debug_problems.code AS dpcode,
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

func (m *Module) getSingleProblem(c echo.Context) (err error) {
	user, err := utility.Verify(c, m.env)
	if user == nil {
		return err
	}

	code := c.Param("code")
	m.env.Log.Infof("Getting problem (%v)", code)

	m.env.Log.Debug("Querying DB for problem")
	var p models.DebugProblem
	err = m.env.DB.Get(&p, getSingleProblemQuery, user.ID, code)

	m.env.Log.Debug("Converting score")
	if !p.RawBestScore.Valid {
		p.BestScore = 0.0
	} else {
		p.BestScore = p.RawBestScore.Float64
	}

	if err != nil {
		if err == sql.ErrNoRows {
			m.env.Log.Infof("Getting problem: (%v) not found", code)
			return c.NoContent(http.StatusNotFound)
		} else {
			m.env.Log.Errorf("Getting problem: Error encountered: %v", err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	m.env.Log.Infof("Found problem (%v)", code)
	return c.JSON(http.StatusOK, p)
}
