package productrepo

import (
	"context"

	"github.com/afif0808/inventory-sales-service/model"
)

type CreateFunc func(ctx context.Context, p model.Product) (model.Product, error)
type GetManyFunc func(ctx context.Context, parameters map[string]interface{}) ([]model.Product, error)
type UpdateManyFunc func(ctx context.Context, p model.Product, parameters map[string]interface{}) error
type DeleteManyFunc func(ctx context.Context, parameters map[string]interface{}) error

func (f GetManyFunc) Generic(ctx context.Context, parameters map[string]interface{}) (interface{}, error) {
	return f(ctx, parameters)
}
