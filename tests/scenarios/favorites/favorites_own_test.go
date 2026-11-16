package favorites

import (
	"fmt"
	"net/http"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/helpers/testdata"
	"api-tests-template/internal/managers/advertisement"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/favorites/models"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestAddOwnAdvertisement() {
	base.Precondition("Покупатель создаёт собственное объявление")
	ownAd := advertisement.CreateAdvertisement(s.T(), s.buyer.Token, testdata.NewAdvertisementRequest(limits.MinPhotosPerAdvertisement))

	base.Step(fmt.Sprintf("Покупатель добавляет собственное объявление %s в избранное", ownAd.Id))
	errorResponse := favorites.AddFavoriteExpectingError(s.T(), s.buyer.Token,
		models.AddFavoriteRequest{AdvertisementId: ownAd.Id}, http.StatusForbidden)

	base.Check("Вернулся 403 и тело ошибки заполнено")
	assertions.ErrorResponseMatches(s.T(), errorResponse, apierrors.FavoriteOwnMessage)

	base.Check("Собственное объявление не попало в избранное покупателя")
	assertions.FavoritesNotContainAdvertisement(s.T(), favorites.GetFavorites(s.T(), s.buyer.Token), ownAd.Id)
}
