package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisements/models"
)

// ContainsAdvertisementByTitle проверяет наличие объявления с указанным названием.
func ContainsAdvertisementByTitle(t *testing.T, items []models.Advertisement, title string) {
	t.Helper()
	for _, item := range items {
		if item.Title == title {
			return
		}
	}
	require.Failf(t, "Объявление не найдено", "В ответе нет объявления с названием %q (найдено объявлений: %d)", title, len(items))
}

// ContainsAdvertisementById проверяет наличие объявления с указанным идентификатором.
func ContainsAdvertisementById(t *testing.T, items []models.Advertisement, id string) {
	t.Helper()
	for _, item := range items {
		if item.Id == id {
			return
		}
	}
	require.Failf(t, "Объявление не найдено", "В ответе нет объявления с id %q (найдено объявлений: %d)", id, len(items))
}
