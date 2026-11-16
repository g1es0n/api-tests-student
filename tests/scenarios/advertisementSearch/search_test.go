package advertisementSearch

import (
	"os"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/helpers/testdata"
	"api-tests-template/internal/managers/advertisement"
	advList "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

const searchSuffixLength = 8

type TestSuite struct {
	suite.Suite
	loginData       models.LoginOkResponse
	existingAdTitle string
	searchBeginning string
	searchMiddle    string
	searchEnd       string
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Авторизация пользователя с кредами из переменных окружения")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

// SetupTest создаёт собственное объявление для каждого поискового сценария.
func (s *TestSuite) SetupTest() {
	base.Precondition("Создаём объявление с уникальными фрагментами названия для поиска")
	// 21 символ в словах и пробелах + 3 * 8 в суффиксах = 45 при лимите API 50.
	suffix := utils.RandomString(searchSuffixLength)
	s.searchBeginning = "Начало" + suffix
	s.searchMiddle = "Середина" + suffix
	s.searchEnd = "Конец" + suffix
	request := testdata.NewAdvertisementRequest(limits.MinPhotosPerAdvertisement)
	request.Title = strings.Join([]string{s.searchBeginning, s.searchMiddle, s.searchEnd}, " ")
	require.LessOrEqual(s.T(), utf8.RuneCountInString(request.Title), limits.MaxTitleLength, "Название фикстуры превышает лимит API")
	created := advertisement.CreateAdvertisement(s.T(), s.loginData.Token, request)
	require.Equal(s.T(), request.Title, created.Title)
	s.existingAdTitle = created.Title
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

func (s *TestSuite) TestSearchByTitle() {
	base.Step("Ищем объявление по фрагменту названия")
	list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, s.searchMiddle)

	s.Run("В ответе есть хотя бы одно объявление", func() {
		require.NotEmpty(s.T(), list.Items)
	})
	s.Run("Среди найденных есть наше объявление", func() {
		assertions.ContainsAdvertisementByTitle(s.T(), list.Items, s.existingAdTitle)
	})
}

func (s *TestSuite) TestSearchByTitle_Substring() {
	cases := []struct {
		name   string
		search string
	}{
		{"начало", s.searchBeginning},
		{"середина", s.searchMiddle},
		{"конец", s.searchEnd},
		{"нижний регистр", strings.ToLower(s.searchMiddle)},
		{"верхний регистр", strings.ToUpper(s.searchMiddle)},
		{"смешанный регистр", utils.MixCase(s.searchMiddle)},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			base.Step("Ищем объявление по фрагменту названия")
			list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, tc.search)
			base.Check("В результатах есть созданное объявление")
			assertions.ContainsAdvertisementByTitle(s.T(), list.Items, s.existingAdTitle)
		})
	}
}
