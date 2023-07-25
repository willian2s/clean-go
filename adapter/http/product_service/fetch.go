package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/willian2s/clean-go/core/dto"
)

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
