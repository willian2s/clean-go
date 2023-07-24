package productusecase

import (
	"github.com/willian2s/clean-go/core/domain"
	"github.com/willian2s/clean-go/core/dto"
)

func (uc usecase) Fetch(paginationRequest *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	products, err := uc.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}
