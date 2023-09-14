package debuguser

import (
	"net/http"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/labstack/echo/v4"
)

const getUserStatsByName = `
SELECT debug_problems.code, MAX(debug_submissions.score) FROM debug_submissions
INNER JOIN teams ON teams.tid = debug_submissions.tid
INNER JOIN debug_problems ON debug_problems.dpid = debug_submissions.dpid
WHERE teams.teamname = ?
GROUP BY debug_submissions.dpid;
`

func (m *Module) GetUserStats(c echo.Context) (err error) {
	userID := c.Param("username")
	m.env.Log.Infof("Getting stats for user (%v)", userID)

	// Query the database
	rows, err := m.env.DB.Query(getUserStatsByName, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}

	result := make(map[string]float64)
	
	// Iterate over the rows
	for rows.Next() {
		var problemID string
		var score float64

		err = rows.Scan(&problemID, &score)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Error: "An error has occured",
			})
		}

		result[problemID] = score
	}

	return c.JSON(http.StatusOK, models.Response{
		Data: result,
	})
}
