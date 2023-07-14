package debugproblem

import (
	"net/http"
	"strconv"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
	"github.com/labstack/echo/v4"
)

// Parts of the query for retrieving all problems from DB.
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
debug_problems.name AS dpname,
debug_problems.language,
debug_problems.result,
score_table.max_score AS best_score

FROM debug_problems
LEFT OUTER JOIN score_table ON score_table.dpid = debug_problems.dpid
`

var getAllProblemQueryFilterRange = `
LIMIT ? OFFSET ?
`

var getAllProblemQueryFilterActive = `
WHERE debug_problems.status = "Active"
`

// Separate queries for admins and non-admins
var getAllProblemQuery = getAllProblemQueryPart1 + getAllProblemQueryFilterActive
var getAllProblemQueryAdmin = getAllProblemQueryPart1

func (m *Module) GetAllProblem(c echo.Context) (err error) {
	// Verify the user first
	user, err := utility.Verify(c, m.env)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}

	// Get the status of the user
	var query string
	if user != nil && user.Status == "Admin" {
		m.env.Log.Info("Getting all problems (Admin)")
		query = getAllProblemQueryAdmin
	} else {
		m.env.Log.Info("Getting all problems")
		query = getAllProblemQuery
	}

	// Check if page_id and limit exists
	pageID_str := c.QueryParam("page_id")
	limit_str := c.QueryParam("limit")

	var pagination bool
	var limit, start int

	if pageID_str != "" && limit_str != "" {
		pagination = true
		query += getAllProblemQueryFilterRange

		pageID, err := strconv.Atoi(pageID_str)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Error: "Cannot convert parameter to int",
			})
		}

		limit, err = strconv.Atoi(limit_str)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Error: "Cannot convert parameter to int",
			})
		}

		start = utility.Pagination(pageID, limit)
	} else {
		pagination = false
	}

	// Query all problems from DB
	listOfProblems := []models.ShortenedProblem{}

	var userID int
	if user == nil {
		userID = 0
	} else {
		userID = user.ID
	}

	if pagination {
		m.env.Log.Debugf("Querying DB for problems from %v (limit %v)", start, limit)
		err = m.env.DB.Select(&listOfProblems, query, userID, limit, start-1)
	} else {
		m.env.Log.Debug("Querying DB for all problems")
		err = m.env.DB.Select(&listOfProblems, query, userID)
	}

	if err != nil {
		m.env.Log.Error("Getting all problems: Error encountered: %v", err)
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}

	// Convert each scores from NULL to 0
	m.env.Log.Debug("Converting score")
	for idx, p := range listOfProblems {
		if user == nil {
			listOfProblems[idx].BestScore = -1.0
		} else {
			if !p.RawBestScore.Valid {
				listOfProblems[idx].BestScore = 0.0
			} else {
				listOfProblems[idx].BestScore = p.RawBestScore.Float64
			}
		}
	}

	// Return all problems
	m.env.Log.Info("Found all problems")
	return c.JSON(http.StatusOK, models.Response{
		Data: listOfProblems,
	})
}
