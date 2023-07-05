package create

import (
	"net/http"
	"strconv"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
	"github.com/labstack/echo/v4"
)

type CheckResult struct {
	Status         string `json:"status"`
	SuggestedCode  string `json:"code,omitempty"`
	DuplicatedCode string `json:"duplicated_code,omitempty"`
	Message        string `json:"message,omitempty"`
}

func (m *Module) Check(c echo.Context) error {
	// Verify the user first
	user, err := utility.Verify(c, m.env)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}
	
	if user == nil {
		return c.JSON(http.StatusForbidden, models.Response{
			Error: "Invalid token",
		})
	}

	if user.Status != "Admin" {
		return c.JSON(http.StatusForbidden, models.Response{
			Error: "You are not allowed to perform this operation",
		})
	}

	m.env.Log.Debugf("ID = %v", c.FormValue("id"))
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Error: "Invalid ID",
		})
	}

	m.env.Log.Infof("Checking information for %v", id)

	duplicated_code, err := checkDuplicated(m.env.DB, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}

	if duplicated_code != "" {
		// DUPLICATED
		m.env.Log.Info("Debug problem with ID %v already existed", id)
		return c.JSON(http.StatusAccepted, models.Response{
			Data: CheckResult{
				Status:         "DUPLICATED",
				DuplicatedCode: duplicated_code,
			},
		})
	}

	// TODO: Check for FAILED

	// Get suggested dpcode
	suggestion, err := getSuggestedCode(m.env.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}

	m.env.Log.Infof("Suggesting %v as code", suggestion)
	return c.JSON(http.StatusAccepted, models.Response{
		Data: CheckResult{
			Status:        "OK",
			SuggestedCode: suggestion,
		},
	})
}
