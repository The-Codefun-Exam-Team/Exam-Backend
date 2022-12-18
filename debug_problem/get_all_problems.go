package debugproblem

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
	"github.com/labstack/echo/v4"
	"net/http"
)

var getAllProblemQueryPart1 = `
WITH score_table AS (
	SELECT 
	
	dpid,
	MAX(score) AS max_score

	FROM debug_submissions

	WHERE tid = ?

	GROUP BY dpid
)

SELECT

debug_problems.code AS dpcode,
MAX(score_table.max_score) AS best_score

FROM debug_problems
LEFT OUTER JOIN debug_submissions ON debug_submissions.dpid = debug_problems.dpid
LEFT OUTER JOIN score_table ON score_table.dpid = debug_problems.dpid
`

var getAllProblemQueryPart2 = `
GROUP BY debug_submissions.dpid
`

var getAllProblemQueryFilterActive = `
	AND debug_problems.status = "Active"
`

var getAllProblemQuery = getAllProblemQueryPart1 + getAllProblemQueryFilterActive + getAllProblemQueryPart2
var getAllProblemQueryAdmin = getAllProblemQueryPart1 + getAllProblemQueryPart2

func (m *Module) getAllProblem(c echo.Context) (err error) {
	user, err := utility.Verify(c, m.env)
	if user == nil {
		return err
	}

	var query string
	if user.Status == "Admin" {
		m.env.Log.Info("Getting all problems (Admin)")
		query = getAllProblemQueryAdmin
	} else {
		m.env.Log.Info("Getting all problems")
		query = getAllProblemQuery
	}

	var listOfProblems []models.ShortenedProblem

	m.env.Log.Debug("Querying DB for all problems")
	err = m.env.DB.Select(&listOfProblems, query, user.ID)

	if err != nil {
		m.env.Log.Error("Getting all problems: Error encountered: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	m.env.Log.Debug("Converting score")
	for idx, p := range listOfProblems {
		if !p.RawBestScore.Valid {
			listOfProblems[idx].BestScore = 0.0
		} else {
			listOfProblems[idx].BestScore = p.RawBestScore.Float64
		}
	}

	m.env.Log.Info("Found all problems")
	return c.JSON(http.StatusOK, listOfProblems)
}