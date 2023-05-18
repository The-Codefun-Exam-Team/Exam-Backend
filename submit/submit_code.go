package submit

import (
	"database/sql"
	"net/http"
	"net/url"
	"strings"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/utility"
	"github.com/labstack/echo/v4"
)

var submitURL = "https://codefun.vn/api/submit"

func (m *Module) SubmitCode(c echo.Context) (err error) {
	user, err := utility.Verify(c, m.env)
	if user == nil {
		return err
	}

	// Get the short name (code) of the problem
	code := c.FormValue("code")
	m.env.Log.Infof("Submitting code for problem %v", code)

	metadata, err := GetMetadata(m.env.DB, code)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no row was found
			m.env.Log.Infof("Getting problem: (%v) not found", code)
			return c.JSON(http.StatusNotFound, models.Response{
				Error: "No problem was found",
			})
		} else {
			// Another error
			m.env.Log.Errorf("Getting problem: Error encountered: %v", err)
			return c.JSON(http.StatusInternalServerError, models.Response{
				Error: "An error has occured",
			})
		}
	}

	// Construct the form to send to codefun.vn
	form := url.Values{}

	form.Add("problem", metadata.PCode)
	form.Add("code", c.FormValue("codetext"))
	form.Add("language", metadata.Language)

	// Construct the request
	request, err := http.NewRequest(
		http.MethodPost, submitURL, strings.NewReader(form.Encode()),
	)

	// Add headers
	utility.AddHeaders(request)
	request.Header.Add("Authorization", c.Request().Header.Get("Authorization"))
	// Add Content-Type header

	// Send the request
	m.env.Log.Debug("Sending request to codefun.vn")
	response, err := m.env.Client.Do(request)
	if err != nil {
		m.env.Log.Errorf("Error sending request: %v", err)
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}

	defer response.Body.Close()

	m.env.Log.Debugf("Status code: %v", response.StatusCode)
	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		m.env.Log.Errorf("Non-2xx status code: %v", err)
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}

	m.env.Log.Debug("Extracting from response")
	rid, err := ExtractResponse(response)
	if err != nil {
		m.env.Log.Errorf("Error extracting from response: %v", err)
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}

	submission := models.DebugSubmission{
		ShortenedSubmission: models.ShortenedSubmission{
			Dpid:       metadata.Dpid,
			Rid:        rid,
			Tid:        user.ID,
			Language:   metadata.Language,
			SubmitTime: utility.GetCurrentTime(),
			Result:     "Q",
			Score:      0,
		},
		Difference: 100000, // Currently hard-coding infinite value
		Code:       c.FormValue("codetext"),
	}

	m.env.Log.Debug("Adding submission to DB")

	id, err := submission.Write(m.env.DB)
	if err != nil {
		m.env.Log.Errorf("Error writing submission to DB: %v", err)
		return c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
	}

	m.env.Log.Infof("Submission %v submitted", id)

	return c.JSON(http.StatusOK, models.Response{
		Data: SubmitReturnValue{
			Drid: id,
		},
	})
}
