package dto

import (
	"net/http"
	"strconv"
	"strings"
)

type PaginationRequestParams struct {
	Search       string   `json:"search"`
	Descending   []string `json:"descending"`
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"itemsPerPage"`
	Sort         []string `json:"sort"`
}

// FromValuePaginationRequestParams generates a PaginationRequestParams object from the given http.Request.
//
// It takes a request object as a parameter and extracts the values of "page" and "itemsPerPage" from the request's form values.
// It then converts these values to integers using the strconv.Atoi() function.
// The extracted values are used to populate the fields of a new PaginationRequestParams object.
//
// Returns a pointer to the populated PaginationRequestParams object and nil error.
func FromValuePaginationRequestParams(request *http.Request) (*PaginationRequestParams, error) {
	page, _ := strconv.Atoi(request.FormValue("page"))
	itemsPerPage, _ := strconv.Atoi(request.FormValue("itemsPerPage"))

	pagiantionRequestParams := PaginationRequestParams{
		Search:       request.FormValue("search"),
		Descending:   strings.Split(request.FormValue("descending"), ","),
		Sort:         strings.Split(request.FormValue("sort"), ","),
		Page:         page,
		ItemsPerPage: itemsPerPage,
	}

	return &pagiantionRequestParams, nil
}
