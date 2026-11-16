package myAdvertisement

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/helpers/users"
	"api-tests-template/internal/managers/auth/models"
	base "api-tests-template/tests"
)

type TestSuite struct {
	suite.Suite
	loginData models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) { suite.Run(t, &TestSuite{}) }

func (s *TestSuite) SetupSuite() { base.SetupSuite() }

// SetupTest изолирует объявления текущего сценария от остальных пакетов и запусков.
func (s *TestSuite) SetupTest() {
	base.Precondition("Регистрируем отдельного владельца объявлений")
	s.loginData = users.NewTestUser(s.T())
}

func (s *TestSuite) TearDownSuite() { base.TearDownSuite() }
