package favorites

import (
	"fmt"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/advertisement"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestDeletedAdvertisementDisappearsFromFavorites() {
	ad := s.createSellerAdvertisement()

	base.Precondition(fmt.Sprintf("Добавляем объявление %s в избранное", ad.Id))
	favorites.AddFavorite(s.T(), s.buyer.Token, ad.Id)
	assertions.FavoritesContainAdvertisement(s.T(), favorites.GetFavorites(s.T(), s.buyer.Token), ad.Id)

	base.Step(fmt.Sprintf("Продавец удаляет объявление %s", ad.Id))
	advertisement.DeleteAdvertisement(s.T(), s.seller.Token, ad.Id)

	base.Check("Объявление исчезло из избранного")
	assertions.FavoritesNotContainAdvertisement(s.T(), favorites.GetFavorites(s.T(), s.buyer.Token), ad.Id)
}
