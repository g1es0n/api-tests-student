package createAdvertisement

import (
	"net/http"

	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/helpers/testdata"
	"api-tests-template/internal/managers/advertisement"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestCreateAdvertisementWithoutAuthorization() {
	for _, tc := range testdata.UnauthorizedCases() {
		s.Run(tc.Name, func() {
			base.Precondition("Готовим валидный запрос на создание объявления")
			request := testdata.NewAdvertisementRequest(limits.MinPhotosPerAdvertisement)

			base.Step("Отправляем запрос на создание объявления " + tc.Name)
			errorResponse := advertisement.CreateAdvertisementExpectingError(s.T(), tc.Token, request, http.StatusUnauthorized)

			base.Check("Вернулся 401 с корректным описанием ошибки")
			tc.AssertError(s.T(), errorResponse)
		})
	}
}
