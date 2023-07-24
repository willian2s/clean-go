package productusecase

import (
	"github.com/willian2s/clean-go/core/domain"
	"github.com/willian2s/clean-go/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParams) (*domain.Pagination, error) {
	products, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}
