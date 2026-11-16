package models

type AdvertisementResponse struct {
	Id          string          `json:"id"`
	UserId      string          `json:"user_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Price       float64         `json:"price"`
	Quantity    int             `json:"quantity"`
	Photos      []PhotoResponse `json:"photos"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}

type PhotoResponse struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	SortOrder int    `json:"sort_order"`
	CreatedAt string `json:"created_at"`
}
