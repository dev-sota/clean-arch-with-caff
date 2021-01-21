package iusecase

import (
	"context"

	"github.com/dev-sota/clean-arch-with-caff/entity"
)

type Administrator interface {
	Create(ctx context.Context, administrator *entity.Administrator) error
	Authentication(ctx context.Context, administrator *entity.Administrator) error
}
