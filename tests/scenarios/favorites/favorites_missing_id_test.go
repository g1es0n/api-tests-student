package favorites

import (
	"net/http"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/favorites/models"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestAddFavoriteWithoutAdvertisementId() {
	base.Step("Отправляем запрос на добавление в избранное с пустым advertisement_id")
	errorResponse := favorites.AddFavoriteExpectingError(s.T(), s.buyer.Token,
		models.AddFavoriteRequest{}, http.StatusBadRequest)

	base.Check("Вернулся 400 и тело ошибки заполнено")
	assertions.ErrorResponseMatches(s.T(), errorResponse, apierrors.FavoriteIDMessage)
}
