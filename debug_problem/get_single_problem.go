package debugproblem

import (
	"database/sql"
	"net/http"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
	"github.com/labstack/echo/v4"
)

type SingleProblem struct {
	BestScore      float64               `json:"best_score"`
	Codetext       string                `json:"code" db:"codetext"`
	Judge          models.Judge          `json:"judge" db:"error"`
	Language       string                `json:"language" db:"language"`
	Result         string                `json:"result" db:"result"`
	CodefunProblem models.CodefunProblem `json:"problem" db:""`
	RawBestScore   sql.NullFloat64       `db:"best_score"`
}

var getSingleProblemQuery = `
SELECT

MAX(debug_submissions.score) AS best_score,
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

func (m *Module) GetSingleProblem(c echo.Context) (err error) {
	user, err := utility.Verify(c, m.env)
	if user == nil {
		return err
	}

	code := c.Param("code")
	m.env.Log.Infof("Getting problem (%v)", code)

	m.env.Log.Debug("Querying DB for problem")
	var p SingleProblem
	err = m.env.DB.QueryRowx(getSingleProblemQuery, user.ID, code).StructScan(&p)

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
			m.env.Log.Error("Getting problem: Error encountered")
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	m.env.Log.Debugf("Getting problem: Found (%v)", code)
	return c.JSON(http.StatusOK, p)
}
