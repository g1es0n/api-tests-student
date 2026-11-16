package favorites

import (
	"net/http"

	"api-tests-template/internal/helpers/testdata"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/favorites/models"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestFavoritesWithoutAuthorization() {
	ad := s.createSellerAdvertisement()

	for _, tc := range testdata.UnauthorizedCases() {
		s.Run("GET /favorites "+tc.Name, func() {
			base.Step("Запрашиваем список избранного " + tc.Name)
			errorResponse := favorites.GetFavoritesExpectingError(s.T(), tc.Token, http.StatusUnauthorized)

			base.Check("Вернулся 401 с корректным описанием ошибки")
			tc.AssertError(s.T(), errorResponse)
		})

		s.Run("POST /favorites "+tc.Name, func() {
			base.Step("Добавляем существующее объявление в избранное " + tc.Name)
			errorResponse := favorites.AddFavoriteExpectingError(s.T(), tc.Token,
				models.AddFavoriteRequest{AdvertisementId: ad.Id}, http.StatusUnauthorized)

			base.Check("Вернулся 401 с корректным описанием ошибки")
			tc.AssertError(s.T(), errorResponse)
		})

		s.Run("DELETE /favorites/{ad_id} "+tc.Name, func() {
			base.Step("Удаляем существующее объявление из избранного " + tc.Name)
			errorResponse := favorites.RemoveFavoriteExpectingError(s.T(), tc.Token, ad.Id, http.StatusUnauthorized)

			base.Check("Вернулся 401 с корректным описанием ошибки")
			tc.AssertError(s.T(), errorResponse)
		})
	}
}
