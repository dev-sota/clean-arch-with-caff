package igateway

import (
	"context"
	"github.com/dev-sota/clean-arch-with-caff/entity"
)

type User interface {
	Get(ctx context.Context, id string) (*entity.User, error)
	List(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}
