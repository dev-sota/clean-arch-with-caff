package igateway

import (
	"context"

	"github.com/dev-sota/clean-arch-with-caff/entity"
)

type Administrator interface {
	Get(ctx context.Context, id string) (*entity.Administrator, error)
	List(ctx context.Context) ([]*entity.Administrator, error)
	Create(ctx context.Context, administrator *entity.Administrator) error
	Update(ctx context.Context, administrator *entity.Administrator) error
	Delete(ctx context.Context, id string) error
}
