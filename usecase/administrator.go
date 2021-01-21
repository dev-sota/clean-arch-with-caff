package usecase

import (
	"context"

	"github.com/dev-sota/clean-arch-with-caff/entity"
	"github.com/dev-sota/clean-arch-with-caff/igateway"
	"github.com/dev-sota/clean-arch-with-caff/iusecase"
	"golang.org/x/crypto/bcrypt"
)

type Administrator struct {
	administratorGateway igateway.Administrator
}

func NewAdministrator(administratorGateway igateway.Administrator) iusecase.Administrator {
	return &Administrator{administratorGateway: administratorGateway}
}

func (a *Administrator) Get(ctx context.Context, id string) (*entity.Administrator, error) {
	return a.administratorGateway.Get(ctx, id)
}

func (a *Administrator) List(ctx context.Context) ([]*entity.Administrator, error) {
	return a.administratorGateway.List(ctx)
}

func (a *Administrator) Create(ctx context.Context, administrator *entity.Administrator) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(administrator.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	administrator.Password = string(hash)
	return a.administratorGateway.Create(ctx, administrator)
}

func (a *Administrator) Update(ctx context.Context, administrator *entity.Administrator) error {
	return a.administratorGateway.Update(ctx, administrator)
}

func (a *Administrator) Delete(ctx context.Context, id string) error {
	return a.administratorGateway.Delete(ctx, id)
}
