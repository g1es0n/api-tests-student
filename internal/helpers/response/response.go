package response

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// ReadBody читает и закрывает тело HTTP-ответа.
func ReadBody(t *testing.T, resp *http.Response) []byte {
	t.Helper()
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Не удалось прочитать тело ответа")

	return body
}

// RequireStatus проверяет HTTP-статус и выводит тело ответа при несовпадении.
func RequireStatus(t *testing.T, resp *http.Response, body []byte, expectedStatus int, action string) {
	t.Helper()
	require.Equalf(t, expectedStatus, resp.StatusCode,
		"%s: ожидался HTTP %d, получен HTTP %d. Тело ответа: %s",
		action, expectedStatus, resp.StatusCode, string(body))
}

// Decode разбирает JSON-ответ в значение указанного типа.
func Decode[T any](t *testing.T, body []byte) T {
	t.Helper()

	var result T
	require.NoErrorf(t, json.Unmarshal(body, &result), "Не удалось разобрать JSON ответа: %s", string(body))

	return result
}
