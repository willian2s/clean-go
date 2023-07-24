package productrepository

import (
	"github.com/willian2s/clean-go/adapter/postgres"
	"github.com/willian2s/clean-go/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
