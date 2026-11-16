package myAdvertisement

import (
	"net/http"

	"api-tests-template/internal/helpers/testdata"
	"api-tests-template/internal/managers/myAdvertisements"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestGetMyAdvertisementsIncorrectToken() {
	for _, tc := range testdata.UnauthorizedCases() {
		s.Run(tc.Name, func() {
			base.Step("Получаем собственные объявления " + tc.Name)
			errorResponse := myAdvertisements.GetMyAdvertisementsExpectingError(s.T(), tc.Token, http.StatusUnauthorized)
			base.Check("Получен 401 с корректной ошибкой авторизации")
			tc.AssertError(s.T(), errorResponse)
		})
	}
}
