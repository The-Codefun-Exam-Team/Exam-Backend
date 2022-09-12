package submit

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/general"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
)

type SubmitResponse struct {
	Rid int `json:"data"`
}

type SubmitReturn struct {
	Drid int `json:"id"`
}

func (g *Group) Submit(c echo.Context) error {
	u, err := models.Verify(c.Request().Header.Get("Authorization"))
	if err != nil {
		return err
		// return c.String(http.StatusOK, fmt.Sprintf("Error while verifying: %v", err))
	}

	if !u.Valid {
		return c.String(http.StatusForbidden, "Invalid token")
	}

	dprob, err := models.ReadDebugProblemWithCode(g.db, c.FormValue("problem"))
	if err != nil {
		return err
		// return c.String(http.StatusOK, fmt.Sprintf("Error while finding dprob (form: %v)): %v", c.FormValue("problem"), err))
	}

	run, err := models.ReadRun(g.db, dprob.Rid)
	if err != nil {
		return err
		// return c.String(http.StatusOK, fmt.Sprintf("Error while finding run: %v", err))
	}

	prob, err := models.ReadProblemWithID(g.db, dprob.Pid)
	if err != nil {
		return err
		// return c.String(http.StatusOK, fmt.Sprintf("Error while finding prob: %v", err))
	}

	form_values_send := url.Values{}

	form_values_send.Add("code", c.FormValue("code"))
	form_values_send.Add("language", run.Language)
	form_values_send.Add("problem", prob.Code)

	req, err := http.NewRequest(http.MethodPost, "https://codefun.vn/api/submit", strings.NewReader(form_values_send.Encode()))
	if err != nil {
		return err
		// return c.String(http.StatusOK, fmt.Sprintf("Error creating request: %v", err))
	}

	req.Header.Add("Authorization", c.Request().Header.Get("Authorization"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Chrome/105.0.0.0")

	rawresp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
		// return c.String(http.StatusOK, fmt.Sprintf("Error processing request: %v", err))
	}

	defer rawresp.Body.Close()

	if rawresp.StatusCode != 201 {
		return c.NoContent(http.StatusBadRequest)
		// return c.String(http.StatusOK, fmt.Sprintf("Status Code: %v", rawresp.StatusCode))
	}

	body, err := io.ReadAll(rawresp.Body)
	if err != nil {
		return err
		// return c.String(http.StatusOK, fmt.Sprintf("Error reading body: %v", err))
	}

	var resp SubmitResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
		// return c.String(http.StatusOK, fmt.Sprintf("Error: %v (Text: %v)", err, body))
	}

	org_code, err := models.ReadSubsCode(g.db, dprob.Rid)
	if err != nil {
		return err
	}

	sub := models.DebugSubmission{
		Dpid:       dprob.Dpid,
		Rid:        resp.Rid,
		Tid:        u.Data.Tid,
		Language:   run.Language,
		Submittime: time.Now().Unix(),
		Result:     "Q",
		Score:      0,
		Diff:       general.EditDistance(c.FormValue("code"), org_code),
		Code:       c.FormValue("code"),
	}

	drid, err := models.WriteDebugSubmission(g.db, &sub)
	if err != nil {
		return err
		// return c.String(http.StatusOK, fmt.Sprintf("Error writing submission: %v", err))
	}

	err = models.AddToQueue(g.db, &models.Queue{
		Rid:  resp.Rid,
		Drid: int(drid),
	})
	if err != nil {
		return err
		// return c.String(http.StatusOK, "Error adding to queue")
	}

	return c.JSON(http.StatusOK, SubmitReturn{
		Drid: int(drid),
	})
}
