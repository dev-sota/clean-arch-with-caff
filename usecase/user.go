package usecase

import (
	"context"
	"github.com/dev-sota/clean-arch-with-caff/igateway"
	"github.com/dev-sota/clean-arch-with-caff/iusecase"
	"github.com/dev-sota/clean-arch-with-caff/entity"
)

type User struct {
	userGateway igateway.User
}

func NewUser(userGateway igateway.User) iusecase.User {
	return &User{ userGateway: userGateway }
}

func (u *User) Get(ctx context.Context, id string) (*entity.User, error) {
	return u.userGateway.Get(ctx, id)
}

func (u *User) List(ctx context.Context) ([]*entity.User, error) {
	return u.userGateway.List(ctx)
}

func (u *User) Create(ctx context.Context, user *entity.User) error {
	return u.userGateway.Create(ctx, user)
}

func (u *User) Update(ctx context.Context, user *entity.User) error {
	return u.userGateway.Update(ctx, user)
}

func (u *User) Delete(ctx context.Context, id string) error {
	return u.userGateway.Delete(ctx, id)
}
