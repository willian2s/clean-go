package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/willian2s/clean-go/core/dto"
)

// @Summary Fetch products with server pagination
// @Description Fetch products with server pagination
// @Tags product
// @Accept  json
// @Produce  json
// @Param sort query string false "1,2"
// @Param descending query string false "true,false"
// @Param page query integer true "1"
// @Param itemsPerPage query integer true "10"
// @Param search query string false "chair"
// @Success 200 {object} domain.Pagination
// @Router /product [get]
func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	pagination, _ := dto.FromValuePaginationRequestParams(request)

	product, err := service.usecase.Fetch(pagination)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(product)
}
