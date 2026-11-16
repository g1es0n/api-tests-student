package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/common/models"
)

// ErrorResponseIsFilled проверяет, что в ответе заполнены поля error и message.
func ErrorResponseIsFilled(t *testing.T, errorResponse models.ErrorResponse) {
	t.Helper()
	require.NotEmpty(t, errorResponse.Error, "Поле error в ответе с ошибкой пустое")
	require.NotEmpty(t, errorResponse.Message, "Поле message в ответе с ошибкой пустое")
}

// UnauthorizedError проверяет тело ответа при запросе без токена.
func UnauthorizedError(t *testing.T, errorResponse models.ErrorResponse) {
	t.Helper()
	require.Equal(t, apierrors.UnauthorizedError, errorResponse.Error, "Некорректное значение поля error")
	require.NotEmpty(t, errorResponse.Message, "Поле message в ответе 401 пустое")
}

// InvalidTokenError проверяет тело ответа при запросе с неверным токеном.
func InvalidTokenError(t *testing.T, errorResponse models.ErrorResponse) {
	t.Helper()
	require.Equal(t, apierrors.UnauthorizedError, errorResponse.Error, "Некорректное значение поля error")
	require.Equal(t, apierrors.InvalidAuthMessage, errorResponse.Message, "Некорректное значение поля message")
}

// ErrorResponseMatches проверяет содержательное сообщение ошибки, а не только непустое тело.
func ErrorResponseMatches(t *testing.T, errorResponse models.ErrorResponse, messagePattern string) {
	t.Helper()
	ErrorResponseIsFilled(t, errorResponse)
	require.Regexp(t, messagePattern, errorResponse.Message,
		"Сообщение не описывает ожидаемую причину ошибки (error=%q)", errorResponse.Error)
}
