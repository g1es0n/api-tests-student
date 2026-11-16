package createAdvertisement

import (
	"fmt"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/helpers/testdata"
	"api-tests-template/internal/managers/advertisement"
	advList "api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestCreateAdvertisementWithAllParameters() {
	base.Precondition(fmt.Sprintf("Готовим объявление с title, description, price, quantity и %d фото",
		limits.MaxPhotosPerAdvertisement))
	request := testdata.NewAdvertisementRequest(limits.MaxPhotosPerAdvertisement)

	base.Step(fmt.Sprintf("Создаём объявление %q", request.Title))
	created := advertisement.CreateAdvertisement(s.T(), s.loginData.Token, request)

	base.Check("Ответ на создание содержит переданные параметры")
	assertions.AdvertisementMatchesRequest(s.T(), created, request)
	require.Equal(s.T(), s.loginData.User.Id, created.UserId, "Объявление создано не от имени текущего пользователя")

	s.Run("GET /advertisement возвращает те же параметры, что передали при создании", func() {
		base.Step(fmt.Sprintf("Получаем объявление %s по id", created.Id))
		fetched := advertisement.GetAdvertisement(s.T(), created.Id)

		base.Check("Параметры объявления совпадают с переданными")
		require.Equal(s.T(), created.Id, fetched.Id, "id объявления не совпадает")
		require.Equal(s.T(), s.loginData.User.Id, fetched.UserId, "user_id не совпадает с текущим пользователем")
		assertions.AdvertisementMatchesRequest(s.T(), fetched, request)
	})

	s.Run("Объявление находится в поиске по полному названию", func() {
		base.Step(fmt.Sprintf("Ищем объявления по названию %q", request.Title))
		list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, request.Title)

		base.Check("Созданное объявление есть в результатах поиска")
		assertions.ContainsAdvertisementById(s.T(), list.Items, created.Id)
	})

	s.Run("Фотографии загружены и доступны через GET /advertisements/{id}/photos", func() {
		base.Step(fmt.Sprintf("Получаем фотографии объявления %s", created.Id))
		photos := advertisement.GetAdvertisementPhotos(s.T(), s.loginData.Token, created.Id)

		base.Check("Фотографии совпадают с загруженными при создании")
		assertions.PhotosMatch(s.T(), photos, created.Photos)
	})
}
