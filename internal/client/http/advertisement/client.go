package advertisement

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/helpers/multipart"
	"api-tests-template/internal/managers/advertisement/models"
)

const (
	formFieldTitle       = "title"
	formFieldDescription = "description"
	formFieldPrice       = "price"
	formFieldQuantity    = "quantity"
	formFieldPhotos      = "photos"
	queryParamId         = "id"
)

func HttpPostAdvertisement(t *testing.T, token string, request models.CreateAdvertisementRequest) *http.Response {
	body, contentType := buildCreateAdvertisementForm(t, request)

	return apiRunner.GetRunner().Auth(token).Create().Post(path.AdvertisementPath).
		ContentType(contentType).
		Body(body).
		Expect(t).
		End().Response
}

func HttpGetAdvertisement(t *testing.T, id string) *http.Response {
	return apiRunner.GetRunner().Create().Get(path.AdvertisementPath).
		Query(queryParamId, id).
		Expect(t).
		End().Response
}

func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.AdvertisementPath).
		Query(queryParamId, id).
		Expect(t).
		End().Response
}

func HttpGetAdvertisementPhotos(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Get(fmt.Sprintf(path.AdvertisementPhotosPathTemplate, url.PathEscape(id))).
		Expect(t).
		End().Response
}

// buildCreateAdvertisementForm собирает поля объявления и фото в тело multipart-запроса.
func buildCreateAdvertisementForm(t *testing.T, request models.CreateAdvertisementRequest) (string, string) {
	form := multipart.NewForm()
	form.AddField(t, formFieldTitle, request.Title)
	form.AddField(t, formFieldDescription, request.Description)

	if request.Price != nil {
		form.AddField(t, formFieldPrice, strconv.FormatFloat(*request.Price, 'f', -1, 64))
	}
	if request.Quantity != nil {
		form.AddField(t, formFieldQuantity, strconv.Itoa(*request.Quantity))
	}
	for _, photo := range request.Photos {
		form.AddFile(t, formFieldPhotos, photo.FileName, photo.ContentType, photo.Content)
	}

	return form.Build(t)
}
