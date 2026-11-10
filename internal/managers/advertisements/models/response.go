package models

type AdvertisementsList struct {
	Items []Advertisement `json:"items"`
}

type Advertisement struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Title  string `json:"title"`
}