package models

type PhotoFile struct {
	FileName    string
	ContentType string
	Content     []byte
}

type CreateAdvertisementRequest struct {
	Title       string
	Description string
	Price       *float64
	Quantity    *int
	Photos      []PhotoFile
}
