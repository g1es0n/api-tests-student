package testdata

import (
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/common/models"
)

type UnauthorizedCase struct {
	Name        string
	Token       string
	AssertError func(t *testing.T, errorResponse models.ErrorResponse)
}

// UnauthorizedCases возвращает варианты запроса без токена и с неверным токеном.
func UnauthorizedCases() []UnauthorizedCase {
	return []UnauthorizedCase{
		{Name: "без токена", Token: "", AssertError: assertions.UnauthorizedError},
		{Name: "с невалидным токеном", Token: apierrors.InvalidToken, AssertError: assertions.InvalidTokenError},
	}
}
