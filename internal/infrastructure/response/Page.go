package response

type Page[T any] struct {
	Total int64 `json:"total"`
	Data  []T   `json:"data"`
}
