package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/favorites/models"
)

// FavoritesContainAdvertisement проверяет наличие объявления в избранном.
func FavoritesContainAdvertisement(t *testing.T, favorites models.FavoriteListResponse, advertisementId string) {
	t.Helper()
	require.Truef(t, containsFavorite(favorites, advertisementId),
		"Объявления %s нет в избранном (всего в избранном: %d)", advertisementId, favorites.Total)
}

// FavoritesNotContainAdvertisement проверяет отсутствие объявления в избранном.
func FavoritesNotContainAdvertisement(t *testing.T, favorites models.FavoriteListResponse, advertisementId string) {
	t.Helper()
	require.Falsef(t, containsFavorite(favorites, advertisementId),
		"Объявление %s всё ещё присутствует в избранном", advertisementId)
}

// containsFavorite ищет объявление в избранном по идентификатору.
func containsFavorite(favorites models.FavoriteListResponse, advertisementId string) bool {
	for _, item := range favorites.Items {
		if item.Advertisement.Id == advertisementId {
			return true
		}
	}
	return false
}
