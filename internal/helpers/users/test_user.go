package users

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/auth"
	authModels "api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"
)

const testUserRandomLength = 24

// NewTestUser регистрирует отдельный аккаунт, чтобы сценарий не зависел от чужих объявлений и избранного.
// API не предоставляет удаление аккаунта; объявления и избранное очищаются менеджерами через t.Cleanup.
func NewTestUser(t *testing.T) authModels.LoginOkResponse {
	t.Helper()
	user := auth.RegisterOrLogin(t, authModels.RegisterRequest{
		Email:     "autotest-" + utils.RandomString(testUserRandomLength) + "@example.com",
		Password:  "At1-" + utils.RandomString(testUserRandomLength),
		FirstName: "Autotest",
		LastName:  "Isolated",
	})
	require.NotEmpty(t, user.Token, "Регистрация не вернула токен")
	require.NotEmpty(t, user.User.Id, "Регистрация не вернула id пользователя")
	return user
}
