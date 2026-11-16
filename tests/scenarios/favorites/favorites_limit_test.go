package favorites

import (
	"fmt"
	"net/http"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/favorites/models"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestFavoritesLimit() {
	base.Precondition("Заполняем избранное отдельного покупателя собственными тестовыми данными")
	var firstID string
	for i := 0; i < limits.MaxFavoritesCount; i++ {
		ad := s.createSellerAdvertisement()
		if i == 0 {
			firstID = ad.Id
		}
		base.Step(fmt.Sprintf("Добавляем объявление %d из %d в избранное", i+1, limits.MaxFavoritesCount))
		favorites.AddFavorite(s.T(), s.buyer.Token, ad.Id)
	}
	extra := s.createSellerAdvertisement()
	base.Step("Добавляем новое чужое объявление сверх лимита избранного")
	errorResponse := favorites.AddFavoriteExpectingError(s.T(), s.buyer.Token,
		models.AddFavoriteRequest{AdvertisementId: extra.Id}, http.StatusBadRequest)
	base.Check("Получен 400 с сообщением о превышении лимита")
	assertions.ErrorResponseMatches(s.T(), errorResponse, apierrors.FavoriteLimitMessage)

	base.Step("Освобождаем одно место и повторно добавляем то же объявление")
	favorites.RemoveFavorite(s.T(), s.buyer.Token, firstID)
	favorites.AddFavorite(s.T(), s.buyer.Token, extra.Id)
	base.Check("После освобождения места добавление успешно")
}
