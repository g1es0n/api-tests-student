package favorites

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/constants/limits"
	"api-tests-template/internal/helpers/testdata"
	"api-tests-template/internal/helpers/users"
	"api-tests-template/internal/managers/advertisement"
	advertisementModels "api-tests-template/internal/managers/advertisement/models"
	authModels "api-tests-template/internal/managers/auth/models"
	base "api-tests-template/tests"
)

type TestSuite struct {
	suite.Suite
	buyer  authModels.LoginOkResponse
	seller authModels.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
	login, password := os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD")

	base.Precondition("Регистрация или авторизация вспомогательного пользователя-продавца")
	s.seller = users.EnsureSeller(s.T(), login, password)
}

// SetupTest создаёт отдельного покупателя с пустым избранным для каждого сценария.
func (s *TestSuite) SetupTest() {
	base.Precondition("Регистрируем отдельного покупателя с пустым избранным")
	s.buyer = users.NewTestUser(s.T())
	s.Require().NotEqual(s.buyer.User.Id, s.seller.User.Id, "Покупатель и продавец должны быть разными пользователями")
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

// createSellerAdvertisement создаёт объявление от имени продавца для текущего теста.
func (s *TestSuite) createSellerAdvertisement() advertisementModels.AdvertisementResponse {
	base.Precondition("Продавец создаёт объявление")
	return advertisement.CreateAdvertisement(s.T(), s.seller.Token,
		testdata.NewAdvertisementRequest(limits.MinPhotosPerAdvertisement))
}
