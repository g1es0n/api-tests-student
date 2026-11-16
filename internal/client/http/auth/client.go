package auth

import (
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	authModels "api-tests-template/internal/managers/auth/models"
)

// HttpPostAuthLogin выполняет POST-запрос для авторизации пользователя
func HttpPostAuthLogin(t *testing.T, request authModels.LoginRequest) *http.Response {
	return apiRunner.GetRunner().Create().Post(path.AuthLoginPath).
		ContentType("application/json").
		JSON(request).
		Expect(t).
		Status(http.StatusOK).
		End().Response
}

func HttpPostAuthRegister(t *testing.T, request authModels.RegisterRequest) *http.Response {
	return apiRunner.GetRunner().Create().Post(path.AuthRegisterPath).
		ContentType("application/json").
		JSON(request).
		Expect(t).
		End().Response
}
