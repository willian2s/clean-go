package productrepository

import (
	"context"

	"github.com/willian2s/clean-go/core/domain"
	"github.com/willian2s/clean-go/core/dto"
	"github.com/willian2s/clean-go/pkg/paginate"
)

// Fetch fetches a paginated list of products from the database.
//
// It takes a PaginationRequestParams pointer as a parameter and returns a Pagination of Product slices and an error.
func (repository repository) Fetch(pagination *dto.PaginationRequestParams) (*domain.Pagination, error) {
	ctx := context.Background()
	products := []domain.Product{}
	total := int(0)

	query, queryCount, err := paginate.Paginate("SELECT * FROM products").
		Page(pagination.Page).
		Desc(pagination.Descending).
		Sort(pagination.Sort).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "name", "description").
		Query()

	if err != nil {
		return nil, err
	}

	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			product := domain.Product{}

			rows.Scan(
				&product.ID,
				&product.Name,
				&product.Price,
				&product.Description,
			)

			products = append(products, product)
		}
	}

	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination{
		Items: products,
		Total: total,
	}, nil
}
