package createAdvertisement

import (
	"net/http"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/helpers/testdata"
	"api-tests-template/internal/managers/advertisement"
	"api-tests-template/internal/managers/advertisement/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

const negativePrice = -100.0

func (s *TestSuite) TestCreateAdvertisementWithInvalidParameters() {
	cases := []struct {
		name           string
		messagePattern string
		modify         func(request *models.CreateAdvertisementRequest)
	}{
		{"пустой title", apierrors.TitleMessage, func(r *models.CreateAdvertisementRequest) { r.Title = "" }},
		{"пустой description", apierrors.DescriptionMessage, func(r *models.CreateAdvertisementRequest) { r.Description = "" }},
		{"отрицательная цена", apierrors.PriceMessage, func(r *models.CreateAdvertisementRequest) { r.Price = utils.Ptr(negativePrice) }},
		{"title длиннее допустимого", apierrors.TitleMessage, func(r *models.CreateAdvertisementRequest) {
			r.Title = utils.RandomString(limits.MaxTitleLength + 1)
		}},
		{"без фотографий", apierrors.PhotosMessage, func(r *models.CreateAdvertisementRequest) { r.Photos = nil }},
		{"фотографий больше допустимого", apierrors.PhotosMessage, func(r *models.CreateAdvertisementRequest) {
			r.Photos = testdata.NewPhotoFiles(limits.MaxPhotosPerAdvertisement + 1)
		}},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			base.Precondition("Готовим валидный запрос и портим его: " + tc.name)
			request := testdata.NewAdvertisementRequest(limits.MinPhotosPerAdvertisement)
			tc.modify(&request)

			base.Step("Отправляем запрос на создание объявления")
			errorResponse := advertisement.CreateAdvertisementExpectingError(s.T(), s.loginData.Token, request, http.StatusBadRequest)

			base.Check("Вернулся 400 с описанием некорректного параметра")
			require.Equal(s.T(), apierrors.InvalidRequest, errorResponse.Error, "Неверный код ошибки валидации")
			assertions.ErrorResponseMatches(s.T(), errorResponse, tc.messagePattern)
		})
	}
}
