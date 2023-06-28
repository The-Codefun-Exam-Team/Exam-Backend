package utility

import (
	"encoding/json"
	"net/http"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/envlib"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
	"github.com/labstack/echo/v4"
)

// Verify check if the user is valid, and if it is, return User.
func Verify(c echo.Context, env *envlib.Env) (*models.User, error) {
	env.Log.Info("Verifying")
	verify, err := VerifyRequest(c, env)
	if err != nil {
		env.Log.Errorf("Verify: Error encountered: %v", err)
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: "An error has occured",
		})
		return nil, err
	}

	if verify.Error != "" {
		env.Log.Info("Verify: Forbidden")
		return nil, c.JSON(http.StatusForbidden, models.Response{
			Error: "Invalid token",
		})
	}

	env.Log.Infof("Verify: User %v", verify.User.ID)
	return &verify.User, nil
}

// VerifyRequest construct, add Authorization and process the http request for verification.
func VerifyRequest(c echo.Context, env *envlib.Env) (*models.ReturnVerify, error) {
	env.Log.Debug("Constructing verify request")
	request, err := ConstructRequest(http.MethodPost, "https://codefun.vn/api/verify")
	if err != nil {
		return nil, err
	}

	token := c.Request().Header.Get("Authorization")
	request.Header.Add("Authorization", token)

	env.Log.Debug("Processing verify request")
	response, err := ProcessRequest(env.Client, request)
	if err != nil {
		return nil, err
	}

	var verify models.ReturnVerify

	env.Log.Debug("Unmarshalling verify response")
	err = json.Unmarshal(response, &verify)
	if err != nil {
		return nil, err
	}

	return &verify, nil
}
