package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisement/models"
)

const priceDelta = 0.001

// AdvertisementMatchesRequest сравнивает поля объявления с данными запроса на создание.
func AdvertisementMatchesRequest(t *testing.T, actual models.AdvertisementResponse, expected models.CreateAdvertisementRequest) {
	t.Helper()
	require.NotEmpty(t, actual.Id, "id объявления пустой")
	require.Equal(t, expected.Title, actual.Title, "title не совпадает с переданным при создании")
	require.Equal(t, expected.Description, actual.Description, "description не совпадает с переданным при создании")
	if expected.Price != nil {
		require.InDelta(t, *expected.Price, actual.Price, priceDelta, "price не совпадает с переданным при создании")
	}
	if expected.Quantity != nil {
		require.Equal(t, *expected.Quantity, actual.Quantity, "quantity не совпадает с переданным при создании")
	}
	require.Lenf(t, actual.Photos, len(expected.Photos), "Количество фото не совпадает с загруженным (%d)", len(expected.Photos))
}

// PhotosMatch сравнивает идентификаторы фотографий и проверяет наличие URL.
func PhotosMatch(t *testing.T, actual []models.PhotoResponse, expected []models.PhotoResponse) {
	t.Helper()
	require.Lenf(t, actual, len(expected), "Количество фотографий не совпадает: ожидалось %d, получено %d", len(expected), len(actual))

	actualIds := make([]string, 0, len(actual))
	for _, photo := range actual {
		require.NotEmptyf(t, photo.Url, "У фотографии %s пустой url", photo.Id)
		actualIds = append(actualIds, photo.Id)
	}

	expectedIds := make([]string, 0, len(expected))
	for _, photo := range expected {
		expectedIds = append(expectedIds, photo.Id)
	}
	require.ElementsMatch(t, expectedIds, actualIds, "id фотографий не совпадают с загруженными при создании")
}
