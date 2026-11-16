package createAdvertisement

import (
	"strings"
	"unicode/utf8"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/helpers/testdata"
	"api-tests-template/internal/managers/advertisement"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestCreateAdvertisementWithMaximumTitle() {
	base.Precondition("Готовим название ровно на границе допустимой длины")
	request := testdata.NewAdvertisementRequest(limits.MinPhotosPerAdvertisement)
	request.Title += strings.Repeat("я", limits.MaxTitleLength-utf8.RuneCountInString(request.Title))
	require.Equal(s.T(), limits.MaxTitleLength, utf8.RuneCountInString(request.Title))
	base.Step("Создаём объявление с максимально допустимым названием")
	created := advertisement.CreateAdvertisement(s.T(), s.loginData.Token, request)
	base.Check("Название и остальные параметры сохранены без изменений")
	assertions.AdvertisementMatchesRequest(s.T(), created, request)
}
