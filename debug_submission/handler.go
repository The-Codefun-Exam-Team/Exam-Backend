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

	u, err := models.Verify(c.Request().Header.Get("Authorization"))
	if err != nil {
		return err
	}

	if !u.Valid {
		return c.String(http.StatusForbidden, "Invalid token")
	}

	models.Resolve1(g.db, id)

	sub, err := models.ReadDebugSubmission(g.db, id)
	if err != nil {
		return err
	}

	if u.Data.Tid != sub.Tid {
		return c.String(http.StatusForbidden, "Not the owner")
	}

	jsub, err := models.ReadJSONDebugSubmission(g.db, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &jsub)
}
