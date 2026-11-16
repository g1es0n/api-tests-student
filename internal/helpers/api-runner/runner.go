package api_runner

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/steinfletcher/apitest"

	"api-tests-template/internal/utils"
)

const (
	defaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36"

	maxRequestAttempts = 5
	retryBaseDelay     = time.Second
	requestTimeout     = 30 * time.Second
)

var httpClient = &http.Client{
	Timeout:   requestTimeout,
	Transport: retryTransport{next: http.DefaultTransport},
}

type ApiTest struct {
	token     string
	userAgent string
	host      string
}

func New() *ApiTest {
	return &ApiTest{
		userAgent: defaultUserAgent,
		host:      os.Getenv("API_URL"),
	}
}

func GetRunner() *ApiTest {
	return New()
}

// Auth Установка Bearer token-а для прохождения авторизации на приватные ручки
func (at *ApiTest) Auth(token string) *ApiTest {
	at.token = token
	return at
}

func (at *ApiTest) Create() *apitest.APITest {
	apitestNew := apitest.New().EnableNetworking(httpClient)
	if os.Getenv("DEBUG") == "true" {
		apitestNew = apitestNew.Debug()
	}
	return apitestNew.
		Intercept(func(request *http.Request) {
			request.Header.Set("User-Agent", at.userAgent)

			if err := mergeServiceUrls(at.host, request.URL); err != nil {
				panic(fmt.Sprintf("Не удалось применить хост из API_URL: %s", err.Error()))
			}

			if len(at.token) != 0 {
				request.Header.Add("Authorization", "Bearer "+at.token)
			}
		})
}

func mergeServiceUrls(host string, r *url.URL) error {
	if host == "" {
		return fmt.Errorf("API_URL не задан: проверьте переменные окружения")
	}

	urlParsed, err := url.Parse(host)
	if err != nil {
		return fmt.Errorf("host cannot be parsed: %s", err.Error())
	}

	if urlParsed.Path != "" {
		r.Path = path.Join(urlParsed.Path, r.Path)
	}

	r.Scheme = urlParsed.Scheme
	r.Host = urlParsed.Host
	return nil
}

type retryTransport struct {
	next http.RoundTripper
}

// RoundTrip отправляет HTTP-запрос и повторяет его при ответах 429 и 503.
func (rt retryTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	var body []byte
	if request.Body != nil {
		var err error
		body, err = io.ReadAll(request.Body)
		_ = request.Body.Close()
		if err != nil {
			return nil, err
		}
	}

	for attempt := 1; ; attempt++ {
		attemptRequest := request.Clone(request.Context())
		if body != nil {
			attemptRequest.Body = io.NopCloser(bytes.NewReader(body))
		}

		resp, err := rt.next.RoundTrip(attemptRequest)
		if err != nil || !isRetryableStatus(resp.StatusCode) || attempt == maxRequestAttempts {
			return resp, err
		}
		_ = resp.Body.Close()

		delay := retryBaseDelay * time.Duration(attempt)
		utils.LogWithLabelAndTimeStamp("Retry", fmt.Sprintf("%s %s: HTTP %d, повтор через %s (попытка %d из %d)",
			request.Method, request.URL.Path, resp.StatusCode, delay, attempt+1, maxRequestAttempts))
		time.Sleep(delay)
	}
}

// isRetryableStatus определяет, следует ли повторить запрос при полученном статусе.
func isRetryableStatus(statusCode int) bool {
	return statusCode == http.StatusServiceUnavailable || statusCode == http.StatusTooManyRequests
}
