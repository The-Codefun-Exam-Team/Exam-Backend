package debugproblem

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
)

func (g *Group) ProblemGet(c echo.Context) error {
	prob, err := models.ReadDebugProblemCode(g.db, c.Param("code"))
	if err != nil {
		return err
	}

	prob.Codetext, err = models.ReadSubsCode(g.db, prob.Rid)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, prob)
}

func testHandler(c echo.Context) error {
	return c.String(http.StatusOK, "problems")
}