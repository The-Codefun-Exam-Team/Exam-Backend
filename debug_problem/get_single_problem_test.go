package debugproblem

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/envlib"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	HTTPFunc func(r *http.Request) (*http.Response, error)
}

func (client *MockClient) Do(r *http.Request) (*http.Response, error)  {
	return client.HTTPFunc(r)
}

func TestGetSingleProblem(t *testing.T) {
	// Setup
	env := envlib.Env{}
	var err error

	env.Log, err = envlib.InitializeLogger("TESTING")
	if err != nil {
		t.Fatal("cannot load logger")
	}

	env.Client = &MockClient{
		HTTPFunc: func(r *http.Request) (*http.Response, error) {
			if r.Header["Authorization"][0] != "Bearer good-token" {
				return &http.Response{
					StatusCode: 403,
					Body: io.NopCloser(strings.NewReader(`{"error": "An error has occured"}`)),
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Body: io.NopCloser(strings.NewReader(`{
					"data": {
					  "id": 1
					}
				  }`)),
			}, nil
		},
	}

	t.Run("D001", func(t *testing.T){
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set("Authorization", "Bearer good-token")

		recorder := httptest.NewRecorder()
		context := echo.New().NewContext(request, recorder)

		context.SetParamNames("code")
		context.SetParamValues("D001")

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal("cannot create a stub database")
		}
		env.DB = sqlx.NewDb(db, "sqlmock")

		m := Module{
			env: &env,
		}

		rows := sqlmock.NewRows([]string{
			"best_score", "dpcode", "dpname", "language", "result", "codetext", "error",
			"pid", "sid", "code", "name", "type", "scoretype", "cid", "status", "pgroup",
			"statement", "timelimit", "score", "usechecker", "checkercode", "solved", "total",
		}).AddRow(
		0, "D001", "D001", "C++", "SS", "<redacted>",
		"2/2////AC|0.001|Accepted||WA|0.001|Wrong Answer", 460, 1, "P148", "P148",
		"", "oi", nil, "Active", "Practice", "BKLR 2019", 1, 100, 0, "", 0, 0)

		mock.ExpectQuery("SELECT").
			WithArgs(1, "D001").
			WillReturnRows(rows)

		if assert.NoError(t, m.GetSingleProblem(context)) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}