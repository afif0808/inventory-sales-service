package productrepo

import (
	"context"

	"github.com/afif0808/inventory-and-sales-service/models"
)

type CreateFunc func(ctx context.Context, p models.Product) (models.Product, error)
type GetManyFunc func(ctx context.Context, parameters map[string]interface{}) ([]models.Product, error)
type UpdateFunc func(ctx context.Context, p models.Product) error
type DeleteMany func(ctx context.Context, parameters map[string]interface{}) error
