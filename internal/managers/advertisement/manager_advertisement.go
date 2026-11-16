package advertisement

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/tidwall/gjson"

	"api-tests-template/internal/client/http/advertisement"
	"api-tests-template/internal/helpers/response"
	"api-tests-template/internal/managers/advertisement/models"
	commonModels "api-tests-template/internal/managers/common/models"
	"api-tests-template/internal/utils"
)

func CreateAdvertisement(t *testing.T, token string, request models.CreateAdvertisementRequest) models.AdvertisementResponse {
	t.Helper()
	raw := advertisement.HttpPostAdvertisement(t, token, request)
	body := response.ReadBody(t, raw)

	if raw.StatusCode == http.StatusCreated {
		registerAdvertisementCleanup(t, token, gjson.GetBytes(body, "id").String())
	}
	response.RequireStatus(t, raw, body, http.StatusCreated, "Создание объявления")

	return response.Decode[models.AdvertisementResponse](t, body)
}

func CreateAdvertisementExpectingError(t *testing.T, token string, request models.CreateAdvertisementRequest,
	expectedStatus int) commonModels.ErrorResponse {
	t.Helper()
	raw := advertisement.HttpPostAdvertisement(t, token, request)
	body := response.ReadBody(t, raw)

	if raw.StatusCode == http.StatusCreated {
		registerAdvertisementCleanup(t, token, gjson.GetBytes(body, "id").String())
	}
	response.RequireStatus(t, raw, body, expectedStatus, "Создание объявления с некорректными данными")

	return response.Decode[commonModels.ErrorResponse](t, body)
}

func GetAdvertisement(t *testing.T, id string) models.AdvertisementResponse {
	t.Helper()
	raw := advertisement.HttpGetAdvertisement(t, id)
	body := response.ReadBody(t, raw)
	response.RequireStatus(t, raw, body, http.StatusOK, fmt.Sprintf("Получение объявления %s", id))

	return response.Decode[models.AdvertisementResponse](t, body)
}

func GetAdvertisementPhotos(t *testing.T, token, id string) []models.PhotoResponse {
	t.Helper()
	raw := advertisement.HttpGetAdvertisementPhotos(t, token, id)
	body := response.ReadBody(t, raw)
	response.RequireStatus(t, raw, body, http.StatusOK, fmt.Sprintf("Получение фотографий объявления %s", id))

	return response.Decode[[]models.PhotoResponse](t, body)
}

func DeleteAdvertisement(t *testing.T, token, id string) {
	t.Helper()
	raw := advertisement.HttpDeleteAdvertisement(t, token, id)
	body := response.ReadBody(t, raw)
	response.RequireStatus(t, raw, body, http.StatusNoContent, fmt.Sprintf("Удаление объявления %s", id))
}

// registerAdvertisementCleanup регистрирует удаление объявления после теста.
func registerAdvertisementCleanup(t *testing.T, token, id string) {
	t.Helper()
	if id == "" {
		return
	}

	t.Cleanup(func() {
		existing := advertisement.HttpGetAdvertisement(t, id)
		existingBody := response.ReadBody(t, existing)
		if existing.StatusCode == http.StatusNotFound {
			utils.LogWithLabelAndTimeStamp("Cleanup", fmt.Sprintf("Объявление %s уже удалено", id))
			return
		}

		utils.LogWithLabelAndTimeStamp("Cleanup", fmt.Sprintf("Удаляем объявление %s", id))
		raw := advertisement.HttpDeleteAdvertisement(t, token, id)
		body := response.ReadBody(t, raw)
		if raw.StatusCode != http.StatusNoContent {
			t.Errorf("Не удалось удалить объявление %s после теста: HTTP %d, тело ответа: %s (GET перед удалением: HTTP %d, %s)",
				id, raw.StatusCode, string(body), existing.StatusCode, string(existingBody))
		}
	})
}
