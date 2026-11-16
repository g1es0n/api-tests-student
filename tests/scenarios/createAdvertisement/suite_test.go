package createAdvertisement

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/helpers/users"
	authModels "api-tests-template/internal/managers/auth/models"
	base "api-tests-template/tests"
)

type TestSuite struct {
	suite.Suite
	loginData authModels.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Регистрация или авторизация пользователя-продавца, от имени которого создаются объявления")
	s.loginData = users.EnsureSeller(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}
