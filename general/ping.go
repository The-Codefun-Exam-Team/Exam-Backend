package general

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pingmongus\n")
}

func TempSubmission(c echo.Context) error {
	id := c.Param("id")
	if id == "1" {
		return c.JSONBlob(http.StatusOK, []byte(`{"codefun_id":2174432,"edit_result":"SS","edit_score":30}`))
	}
	if id == "2" {
		return c.JSONBlob(http.StatusOK, []byte(`{"codefun_id":2174432,"edit_result":"AC","edit_score":100}`))
	}
	return c.JSONBlob(http.StatusOK, []byte(`{"codefun_id":2174432,"edit_result":"Q","edit_score":0}`))
}
