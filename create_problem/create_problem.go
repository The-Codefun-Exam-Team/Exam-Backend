package create

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
	"github.com/labstack/echo/v4"
)

type CreateResult struct {
	Status string `json:"status"`
	Message string `json:"message"`
	DPCode string `json:"code,omitempty"`
}

func (m *Module) CreateProblem(c echo.Context) (err error) {
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

	m.env.Log.Infof("(form) id is %v", c.FormValue("id"))
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Error: "Invalid ID",
		})
	}
	m.env.Log.Infof("Creating problem for submission %v", id)

	m.env.Log.Debugf("Checking for duplication")
	
	duplicated_code, err := checkDuplicated(m.env.DB, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}

	if duplicated_code != "" {
		m.env.Log.Infof("Problem already existed for ID %v (%v)", id, duplicated_code)
		return c.JSON(http.StatusAccepted, models.Response{
			Data: CreateResult{
				Status: "DUPLICATED",
				Message: fmt.Sprintf("Problem already existed (%v)", duplicated_code),
			},
		})
	}

	m.env.Log.Debugf("Getting problem code")

	dpcode := c.FormValue("code")
	if dpcode == "" {
		m.env.Log.Debugf("Getting problem code suggestion")
		dpcode, err = getSuggestedCode(m.env.DB)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Error: "An error has occured",
			})
		}
		m.env.Log.Infof("Suggesting %v as problem code", dpcode)
	}

	m.env.Log.Debugf("Getting problem name")

	name := c.FormValue("name")
	if name == "" {
		name = dpcode
	}

	// TODO: Check for FAILED

	// TODO: Actually adding the problem

	// Get run data
	// Create problem
	// Write to DB

	m.env.Log.Infof("Created problem %v", dpcode)
	return c.JSON(http.StatusAccepted, models.Response{
		Data: CreateResult{
			Status: "OK",
			Message: "Problem added successfully",
			DPCode: dpcode,
		},
	})
}
