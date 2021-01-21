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

func (a *Administrator) Create(ctx context.Context, administrator *entity.Administrator) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(administrator.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	administrator.Password = string(hash)
	return a.administratorGateway.Create(ctx, administrator)
}

func (a *Administrator) Authentication(ctx context.Context, administrator *entity.Administrator) error {
	res, err := a.administratorGateway.GetByEmail(ctx, administrator.Email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(administrator.Password))
	return err
}
