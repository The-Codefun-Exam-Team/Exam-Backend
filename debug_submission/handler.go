package debugsubmission

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
)

func (g *Group) SubmissionGet(c echo.Context) error {
	log.Print("Getting ID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	log.Print("Verifying")

	u, err := models.Verify(c.Request().Header.Get("Authorization"))
	if err != nil {
		return err
	}

	if !u.Valid {
		return c.String(http.StatusForbidden, "Invalid token")
	}

	log.Print("Reading debug submission")

	sub, err := models.ReadDebugSubmission(g.db, id)
	if err != nil {
		return err
	}

	log.Print("Checking owner")

	if u.Data.Tid != sub.Tid {
		return c.String(http.StatusForbidden, "Not the owner")
	}

	log.Print("Reading JSON Submission")

	jsub, err := models.ReadJSONDebugSubmission(g.db, id)
	if err != nil {
		return err
	}

	log.Print("Returning")

	return c.JSON(http.StatusOK, &jsub)
}
