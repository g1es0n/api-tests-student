package myAdvertisement

import (
	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/helpers/testdata"
	"api-tests-template/internal/managers/advertisement"
	advertisementModels "api-tests-template/internal/managers/advertisement/models"
	"api-tests-template/internal/managers/myAdvertisements"
	base "api-tests-template/tests"
)

const ownAdvertisementsCount = 2

func (s *TestSuite) TestGetMyAdvertisementsPositive() {
	expected := make(map[string]advertisementModels.CreateAdvertisementRequest, ownAdvertisementsCount)
	base.Precondition("Создаём два собственных объявления с фотографиями")
	for i := 0; i < ownAdvertisementsCount; i++ {
		request := testdata.NewAdvertisementRequest(limits.MinPhotosPerAdvertisement)
		created := advertisement.CreateAdvertisement(s.T(), s.loginData.Token, request)
		expected[created.Id] = request
	}

	base.Step("Получаем список объявлений отдельного владельца")
	list := myAdvertisements.GetMyAdvertisements(s.T(), s.loginData.Token)

	base.Check("В списке ровно созданные объявления, параметры и фотографии совпадают")
	require.Len(s.T(), list.Items, ownAdvertisementsCount)
	seen := make(map[string]bool, ownAdvertisementsCount)
	for _, item := range list.Items {
		request, exists := expected[item.Id]
		require.Truef(s.T(), exists, "В списке неожиданное объявление %s", item.Id)
		require.Falsef(s.T(), seen[item.Id], "Объявление %s продублировано", item.Id)
		seen[item.Id] = true
		require.Equalf(s.T(), s.loginData.User.Id, item.UserId, "Неверный владелец объявления %s", item.Id)
		require.NotEmptyf(s.T(), item.CreatedAt, "created_at пустое у %s", item.Id)
		require.NotEmptyf(s.T(), item.UpdatedAt, "updated_at пустое у %s", item.Id)
		assertions.AdvertisementMatchesRequest(s.T(), item, request)
	}
}
