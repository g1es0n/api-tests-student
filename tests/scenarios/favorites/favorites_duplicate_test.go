package favorites

import (
	"fmt"
	"net/http"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/favorites/models"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestAddSameAdvertisementTwice() {
	ad := s.createSellerAdvertisement()

	base.Precondition(fmt.Sprintf("Добавляем объявление %s в избранное", ad.Id))
	favorites.AddFavorite(s.T(), s.buyer.Token, ad.Id)

	base.Step("Повторно добавляем то же объявление в избранное")
	errorResponse := favorites.AddFavoriteExpectingError(s.T(), s.buyer.Token,
		models.AddFavoriteRequest{AdvertisementId: ad.Id}, http.StatusConflict)

	base.Check("Вернулся 409 и тело ошибки заполнено")
	assertions.ErrorResponseMatches(s.T(), errorResponse, apierrors.FavoriteDuplicateMessage)
}
