package create

import (
	"net/http"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
	"github.com/labstack/echo/v4"
)

func (m *Module) CreateProblem(c echo.Context) (err error) {
	// Verify the user first
	user, err := utility.Verify(c, m.env)
	if user == nil {
		return err
	}

	if user.Status != "Admin" {
		return c.JSON(http.StatusForbidden, models.Response{
			Error: "You are not allowed to perform this operation",
		})
	}

	// TODO(unknown): Create the problem

	return nil
}