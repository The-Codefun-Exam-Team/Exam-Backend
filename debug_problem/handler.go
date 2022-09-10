package debugproblem

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
)

func (g *Group) ProblemGet(c echo.Context) error {
	prob, err := models.ReadJSONDebugProblemWithCode(g.db, c.Param("code"))
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: '%v'", err))
	}

	return c.JSON(http.StatusOK, prob)
}

func testHandler(c echo.Context) error {
	return c.String(http.StatusOK, "problems")
}
