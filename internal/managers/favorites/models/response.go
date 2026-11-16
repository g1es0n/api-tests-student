package models

import advertisementModels "api-tests-template/internal/managers/advertisement/models"

type FavoriteResponse struct {
	FavoriteId    string                                    `json:"favorite_id"`
	AddedAt       string                                    `json:"added_at"`
	Advertisement advertisementModels.AdvertisementResponse `json:"advertisement"`
}

type FavoriteListResponse struct {
	Items      []FavoriteResponse `json:"items"`
	Page       int                `json:"page"`
	PageSize   int                `json:"page_size"`
	Total      int                `json:"total"`
	TotalPages int                `json:"total_pages"`
}
