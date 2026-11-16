package apierrors

// Шаблоны проверяют причину отказа, не привязываясь к пунктуации сообщения.
const (
	InvalidRequest           = "invalid_request"
	TitleMessage             = `(?i)title`
	DescriptionMessage       = `(?i)description`
	PriceMessage             = `(?i)price`
	PhotosMessage            = `(?i)photo`
	FavoriteDuplicateMessage = `(?i)already`
	FavoriteOwnMessage       = `(?i)own`
	FavoriteIDMessage        = `^invalid UUID length: 0$`
	NotFoundMessage          = `(?i)(not found|does not exist)`
	FavoriteLimitMessage     = `(?i)(limit|maximum|max|50)`
)
