package general

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pingmongus\n")
}

func TempDebug(c echo.Context) error {
	json_str := `{ "data":{ "problem":{ "code": "P001", "id": 23, "name": "P001" }, "language": "C++", "result": "AC", "score": 100, "code": "#include <iostream>\nusing namespace std;\n\nint main() {\n int x;\n cin >> x;\n cout << 2*x;\n}" } }`
	return c.JSONBlob(http.StatusOK, []byte(json_str))
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
