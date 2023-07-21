package domain

type Pagination[T any] struct {
	Items T   `json:"items"`
	Total int `json:"total"`
}
