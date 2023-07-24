package domain

type Pagination struct {
	Items interface{} `json:"items"`
	Total int         `json:"total"`
}
