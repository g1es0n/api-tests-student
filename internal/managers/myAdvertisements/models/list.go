package models

import advertisementModels "api-tests-template/internal/managers/advertisement/models"

type AdvertisementsList struct {
	Items      []advertisementModels.AdvertisementResponse `json:"items"`
	Total      int                                         `json:"total"`
	HasNext    bool                                        `json:"has_next"`
	NextCursor string                                      `json:"next_cursor"`
}
