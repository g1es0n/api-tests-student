package favorites

import (
	"fmt"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestAddAndRemoveFavorite() {
	ad := s.createSellerAdvertisement()

	base.Step(fmt.Sprintf("Добавляем объявление %s в избранное", ad.Id))
	added := favorites.AddFavorite(s.T(), s.buyer.Token, ad.Id)

	base.Check("Ответ на добавление содержит нужное объявление")
	require.NotEmpty(s.T(), added.FavoriteId, "favorite_id пустой")
	require.Equal(s.T(), ad.Id, added.Advertisement.Id, "В ответе на добавление другое объявление")

	base.Check("Объявление появилось в списке избранного")
	assertions.FavoritesContainAdvertisement(s.T(), favorites.GetFavorites(s.T(), s.buyer.Token), ad.Id)

	base.Step(fmt.Sprintf("Удаляем объявление %s из избранного", ad.Id))
	favorites.RemoveFavorite(s.T(), s.buyer.Token, ad.Id)

	base.Check("Объявление пропало из списка избранного")
	assertions.FavoritesNotContainAdvertisement(s.T(), favorites.GetFavorites(s.T(), s.buyer.Token), ad.Id)
}
