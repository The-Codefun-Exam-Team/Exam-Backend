package debugproblem

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
	testcases := []struct{
		code string
		token string
		errorcode int
		rows *sqlmock.Rows
	}{
		{
			code: "D001",
			token: "good-token",
			errorcode: http.StatusOK,
			rows: sqlmock.NewRows([]string{
					"best_score", "dpcode", "dpname", "language", "result", "codetext", "error",
					"pid", "sid", "code", "name", "type", "scoretype", "cid", "status", "pgroup",
					"statement", "timelimit", "score", "usechecker", "checkercode", "solved", "total",
				}).AddRow(
				0, "D001", "D001", "C++", "SS", "<redacted>",
				"2/2////AC|0.001|Accepted||WA|0.001|Wrong Answer", 460, 1, "P148", "P148",
				"", "oi", nil, "Active", "Practice", "BKLR 2019", 1, 100, 0, "", 0, 0),
		},
		{
			code: "D001",
			token: "bad-token",
			errorcode: http.StatusForbidden,
			rows: sqlmock.NewRows([]string{
					"best_score", "dpcode", "dpname", "language", "result", "codetext", "error",
					"pid", "sid", "code", "name", "type", "scoretype", "cid", "status", "pgroup",
					"statement", "timelimit", "score", "usechecker", "checkercode", "solved", "total",
				}).AddRow(
				0, "D001", "D001", "C++", "SS", "<redacted>",
				"2/2////AC|0.001|Accepted||WA|0.001|Wrong Answer", 460, 1, "P148", "P148",
				"", "oi", nil, "Active", "Practice", "BKLR 2019", 1, 100, 0, "", 0, 0),
		},
		{
			code: "non-existent-problem",
			token: "good-token",
			errorcode: http.StatusNotFound,
			rows: sqlmock.NewRows([]string{
					"best_score", "dpcode", "dpname", "language", "result", "codetext", "error",
					"pid", "sid", "code", "name", "type", "scoretype", "cid", "status", "pgroup",
					"statement", "timelimit", "score", "usechecker", "checkercode", "solved", "total",
				}),
		},
	}

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
					  "id": 1234
					}
				  }`)),
			}, nil
		},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%s %s", test.code, test.token), func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/", nil)
			request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", test.token))

			recorder := httptest.NewRecorder()
			context := echo.New().NewContext(request, recorder)

			context.SetParamNames("code")
			context.SetParamValues(test.code)

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal("cannot create a stub database")
			}
			env.DB = sqlx.NewDb(db, "sqlmock")
			
			m := Module{
				env: &env,
			}

			if test.token == "good-token" {
				mock.ExpectQuery("SELECT").
					WithArgs(1234, test.code). // 1234 is the ID for the mock user
					WillReturnRows(test.rows)
			}

			if assert.NoError(t, m.GetSingleProblem(context), "Error encountered") {
				assert.Equal(t, test.errorcode, recorder.Code, "Wrong status code")
			}
	
			assert.NoError(t, mock.ExpectationsWereMet(), "SQL query not matching")
		})
	}
}