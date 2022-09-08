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