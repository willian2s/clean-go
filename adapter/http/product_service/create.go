package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/willian2s/clean-go/core/dto"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	productRequest, err := dto.FromJSONCreateProductRequest(request.Body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	product, err := service.usecase.Create(productRequest)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(product)
}
