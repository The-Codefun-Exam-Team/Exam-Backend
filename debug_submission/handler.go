package debugsubmission

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
)

func (g *Group) SubmissionGet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	sub, err := models.ReadJSONDebugSubmission(g.db, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &sub)
}
