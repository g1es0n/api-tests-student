package favorites

import (
	"fmt"
	"net/http"
	"testing"

	"api-tests-template/internal/client/http/favorites"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/helpers/response"
	commonModels "api-tests-template/internal/managers/common/models"
	"api-tests-template/internal/managers/favorites/models"
	"api-tests-template/internal/utils"
)

func GetFavorites(t *testing.T, token string) models.FavoriteListResponse {
	t.Helper()
	raw := favorites.HttpGetFavorites(t, token, limits.MaxFavoritesPageSize)
	body := response.ReadBody(t, raw)
	response.RequireStatus(t, raw, body, http.StatusOK, "Получение списка избранного")

	return response.Decode[models.FavoriteListResponse](t, body)
}

func GetFavoritesExpectingError(t *testing.T, token string, expectedStatus int) commonModels.ErrorResponse {
	t.Helper()
	raw := favorites.HttpGetFavorites(t, token, limits.MaxFavoritesPageSize)
	body := response.ReadBody(t, raw)
	response.RequireStatus(t, raw, body, expectedStatus, "Получение списка избранного")

	return response.Decode[commonModels.ErrorResponse](t, body)
}

func AddFavorite(t *testing.T, token, advertisementId string) models.FavoriteResponse {
	t.Helper()
	raw := favorites.HttpPostFavorite(t, token, models.AddFavoriteRequest{AdvertisementId: advertisementId})
	body := response.ReadBody(t, raw)

	if raw.StatusCode == http.StatusCreated {
		registerFavoriteCleanup(t, token, advertisementId)
	}
	response.RequireStatus(t, raw, body, http.StatusCreated, fmt.Sprintf("Добавление объявления %s в избранное", advertisementId))

	return response.Decode[models.FavoriteResponse](t, body)
}

func AddFavoriteExpectingError(t *testing.T, token string, request models.AddFavoriteRequest,
	expectedStatus int) commonModels.ErrorResponse {
	t.Helper()
	raw := favorites.HttpPostFavorite(t, token, request)
	body := response.ReadBody(t, raw)

	if raw.StatusCode == http.StatusCreated {
		registerFavoriteCleanup(t, token, request.AdvertisementId)
	}
	response.RequireStatus(t, raw, body, expectedStatus,
		fmt.Sprintf("Добавление объявления %q в избранное", request.AdvertisementId))

	return response.Decode[commonModels.ErrorResponse](t, body)
}

func RemoveFavorite(t *testing.T, token, advertisementId string) {
	t.Helper()
	raw := favorites.HttpDeleteFavorite(t, token, advertisementId)
	body := response.ReadBody(t, raw)
	response.RequireStatus(t, raw, body, http.StatusNoContent, fmt.Sprintf("Удаление объявления %s из избранного", advertisementId))
}

func RemoveFavoriteExpectingError(t *testing.T, token, advertisementId string, expectedStatus int) commonModels.ErrorResponse {
	t.Helper()
	raw := favorites.HttpDeleteFavorite(t, token, advertisementId)
	body := response.ReadBody(t, raw)
	response.RequireStatus(t, raw, body, expectedStatus, fmt.Sprintf("Удаление объявления %q из избранного", advertisementId))

	return response.Decode[commonModels.ErrorResponse](t, body)
}

// registerFavoriteCleanup регистрирует удаление объявления из избранного после теста.
func registerFavoriteCleanup(t *testing.T, token, advertisementId string) {
	t.Helper()
	if advertisementId == "" {
		return
	}

	t.Cleanup(func() {
		utils.LogWithLabelAndTimeStamp("Cleanup", fmt.Sprintf("Удаляем объявление %s из избранного", advertisementId))
		raw := favorites.HttpDeleteFavorite(t, token, advertisementId)
		body := response.ReadBody(t, raw)
		if raw.StatusCode != http.StatusNoContent && raw.StatusCode != http.StatusNotFound {
			t.Errorf("Не удалось удалить объявление %s из избранного после теста: HTTP %d, тело ответа: %s",
				advertisementId, raw.StatusCode, string(body))
		}
	})
}
