package users

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"api-tests-template/internal/managers/auth"
	authModels "api-tests-template/internal/managers/auth/models"
)

const (
	sellerFirstName    = "Seller"
	sellerLastName     = "Autotest"
	sellerEmailDomain  = "example.com"
	sellerEmailHashLen = 12
	sellerPasswordLen  = 20
)

// EnsureSeller возвращает данные авторизации тестового пользователя-продавца.
func EnsureSeller(t *testing.T, mainLogin, mainPassword string) authModels.LoginOkResponse {
	t.Helper()

	email, password := os.Getenv("TEST_SELLER_LOGIN"), os.Getenv("TEST_SELLER_PASSWORD")
	if email == "" || password == "" {
		email, password = deriveSellerCredentials(mainLogin, mainPassword)
	}

	return auth.RegisterOrLogin(t, authModels.RegisterRequest{
		Email:     email,
		Password:  password,
		FirstName: sellerFirstName,
		LastName:  sellerLastName,
	})
}

// deriveSellerCredentials вычисляет стабильные логин и пароль продавца из данных основного аккаунта.
func deriveSellerCredentials(mainLogin, mainPassword string) (string, string) {
	loginHash := sha256.Sum256([]byte(mainLogin))
	passwordHash := sha256.Sum256([]byte(mainLogin + ":" + mainPassword))

	email := fmt.Sprintf("seller-%s@%s", hex.EncodeToString(loginHash[:])[:sellerEmailHashLen], sellerEmailDomain)
	password := "Sl-" + hex.EncodeToString(passwordHash[:])[:sellerPasswordLen]

	return email, password
}
