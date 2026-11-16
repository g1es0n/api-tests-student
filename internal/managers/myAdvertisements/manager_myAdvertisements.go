package myAdvertisements

import (
	"net/http"
	"testing"

	"api-tests-template/internal/client/http/myAdvertisements"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/helpers/response"
	commonModels "api-tests-template/internal/managers/common/models"
	"api-tests-template/internal/managers/myAdvertisements/models"
)

// GetMyAdvertisements получает страницу собственных объявлений и проверяет HTTP 200.
func GetMyAdvertisements(t *testing.T, token string) models.AdvertisementsList {
	t.Helper()
	raw := myAdvertisements.HttpGetMyAdvertisements(t, token, limits.MyAdvertisementsPageSize)
	body := response.ReadBody(t, raw)
	response.RequireStatus(t, raw, body, http.StatusOK, "Получение собственных объявлений")
	return response.Decode[models.AdvertisementsList](t, body)
}

// GetMyAdvertisementsExpectingError проверяет ожидаемый статус и разбирает ошибку списка объявлений.
func GetMyAdvertisementsExpectingError(t *testing.T, token string, expectedStatus int) commonModels.ErrorResponse {
	t.Helper()
	raw := myAdvertisements.HttpGetMyAdvertisements(t, token, limits.MyAdvertisementsPageSize)
	body := response.ReadBody(t, raw)
	response.RequireStatus(t, raw, body, expectedStatus, "Получение собственных объявлений без авторизации")
	return response.Decode[commonModels.ErrorResponse](t, body)
}
