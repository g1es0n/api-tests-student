package testdata

import (
	"fmt"

	"api-tests-template/internal/managers/advertisement/models"
	"api-tests-template/internal/utils"
)

const (
	titlePrefix        = "Автотест объявление"
	description        = "Объявление создано автотестом и будет удалено после прогона"
	price              = 4990.5
	quantity           = 3
	randomSuffixLength = 12
	photoSize          = 64
	photoContentType   = "image/png"
)

// NewAdvertisementRequest создаёт валидный запрос с уникальным названием и заданным числом фото.
func NewAdvertisementRequest(photosCount int) models.CreateAdvertisementRequest {
	return models.CreateAdvertisementRequest{
		Title:       fmt.Sprintf("%s %s", titlePrefix, utils.RandomString(randomSuffixLength)),
		Description: description,
		Price:       utils.Ptr(price),
		Quantity:    utils.Ptr(quantity),
		Photos:      NewPhotoFiles(photosCount),
	}
}

// NewPhotoFiles генерирует заданное количество файлов PNG для загрузки.
func NewPhotoFiles(count int) []models.PhotoFile {
	photos := make([]models.PhotoFile, 0, count)
	for i := 1; i <= count; i++ {
		photos = append(photos, models.PhotoFile{
			FileName:    fmt.Sprintf("photo_%d.png", i),
			ContentType: photoContentType,
			Content:     utils.GeneratePNG(photoSize, photoSize),
		})
	}

	return photos
}
