package debugproblem

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
)

func (g *Group) ProblemGet(c echo.Context) error {
	u, err := models.Verify(c.Request().Header.Get("Authorization"))
	if err != nil {
		// return err
		return c.String(http.StatusOK, fmt.Sprintf("Error while verifying: %v", err))
	}

	if !u.Valid {
		return c.String(http.StatusForbidden, "Invalid token")
	}

	prob, err := models.ReadJSONDebugProblemWithCode(g.db, c.Param("code"), u.Data.Tid)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: '%v'", err))
	}

	return c.JSON(http.StatusOK, prob)
}

func testHandler(c echo.Context) error {
	return c.String(http.StatusOK, "problems")
}
