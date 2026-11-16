package favorites

import (
	"net/http"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/favorites/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestAddNotExistingAdvertisementToFavorites() {
	notExistingId := utils.RandomUUID()

	base.Step("Добавляем в избранное несуществующее объявление " + notExistingId)
	errorResponse := favorites.AddFavoriteExpectingError(s.T(), s.buyer.Token,
		models.AddFavoriteRequest{AdvertisementId: notExistingId}, http.StatusNotFound)

	base.Check("Вернулся 404 и тело ошибки заполнено")
	assertions.ErrorResponseMatches(s.T(), errorResponse, apierrors.NotFoundMessage)
}
