package favorites

import (
	"net/http"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestRemoveNotExistingAdvertisementFromFavorites() {
	notExistingId := utils.RandomUUID()

	base.Step("Удаляем из избранного несуществующее объявление " + notExistingId)
	errorResponse := favorites.RemoveFavoriteExpectingError(s.T(), s.buyer.Token, notExistingId, http.StatusNotFound)

	base.Check("Вернулся 404 и тело ошибки заполнено")
	assertions.ErrorResponseMatches(s.T(), errorResponse, apierrors.NotFoundMessage)
}
