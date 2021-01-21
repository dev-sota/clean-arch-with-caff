package igateway

import (
	"context"

	"github.com/dev-sota/clean-arch-with-caff/entity"
)

type Administrator interface {
	Create(ctx context.Context, administrator *entity.Administrator) error
	GetByEmail(ctx context.Context, email string) (*entity.Administrator, error)
}
