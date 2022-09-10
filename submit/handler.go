package submit

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
)

type SubmitResponse struct {
	Rid int `json:"data"`
}

type SubmitReturn struct {
	Drid int `json:"id"`
}

func (g *Group) Submit(c echo.Context) error {
	form_values_get, err := c.FormParams()
	if err != nil {
		// return err
		return c.String(http.StatusOK, fmt.Sprintf("Error: %v", err))
	}

	u, err := models.Verify(g.db, c.Request().Header.Get("Authorization"))
	if err != nil {
		// return err
		return c.String(http.StatusOK, fmt.Sprintf("Error: %v", err))
	}

	if !u.Valid {
		return c.String(http.StatusOK, "Invalid token")
	}

	dprob, err := models.ReadDebugProblemWithCode(g.db, form_values_get.Get("problem"))
	if err != nil {
		// return err
		return c.String(http.StatusOK, fmt.Sprintf("Error: %v", err))
	}

	run, err := models.ReadRun(g.db, dprob.Rid)
	if err != nil {
		// return err
		return c.String(http.StatusOK, fmt.Sprintf("Error: %v", err))
	}

	prob, err := models.ReadProblemWithID(g.db, dprob.Pid)
	if err != nil {
		// return err
		return c.String(http.StatusOK, fmt.Sprintf("Error: %v", err))
	}

	form_values_send := url.Values{}

	form_values_send.Add("code", form_values_get.Get("code"))
	form_values_send.Add("language", run.Language)
	form_values_send.Add("problem", prob.Code)

	req, err := http.NewRequest(http.MethodGet, "https://codefun.vn/api/submit", strings.NewReader(form_values_send.Encode()))
	if err != nil {
		// return err
		return c.String(http.StatusOK, fmt.Sprintf("Error: %v", err))
	}

	req.Header.Add("Authorization", c.Request().Header.Get("Authorization"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Chrome/105.0.0.0")

	rawresp, err := http.DefaultClient.Do(req)
	if err != nil {
		// return err
		return c.String(http.StatusOK, fmt.Sprintf("Error: %v", err))
	}

	defer rawresp.Body.Close()

	body, err := io.ReadAll(rawresp.Body)
	if err != nil {
		// return err
		return c.String(http.StatusOK, fmt.Sprintf("Error: %v", err))
	}

	var resp SubmitResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		// return err
		return c.String(http.StatusOK, fmt.Sprintf("Error: %v (Text: %v)", err, body))
	}

	sub := models.DebugSubmission{
		Dpid: dprob.Dpid,
		Tid: u.Data.Tid,
		Language: run.Language,
		Submittime: time.Now().Unix(),
		Score: 100,
		Diff: 0,
	}

	drid, err := models.WriteDebugSubmission(g.db, &sub)
	if err != nil {
		// return err
		return c.String(http.StatusOK, fmt.Sprintf("Error: %v", err))
	}

	return c.JSON(http.StatusOK, SubmitReturn{
		Drid: int(drid),
	})
}