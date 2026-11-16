package favorites

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/favorites/models"
)

func HttpGetFavorites(t *testing.T, token string, pageSize int) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Get(path.FavoritesPath).
		Query("page_size", strconv.Itoa(pageSize)).
		Expect(t).
		End().Response
}

func HttpPostFavorite(t *testing.T, token string, request models.AddFavoriteRequest) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Post(path.FavoritesPath).
		ContentType("application/json").
		JSON(request).
		Expect(t).
		End().Response
}

func HttpDeleteFavorite(t *testing.T, token, advertisementId string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Delete(fmt.Sprintf(path.FavoritePathTemplate, url.PathEscape(advertisementId))).
		Expect(t).
		End().Response
}
