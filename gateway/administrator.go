package gateway

import (
	"context"

	"github.com/dev-sota/clean-arch-with-caff/entity"
	"github.com/dev-sota/clean-arch-with-caff/igateway"
	"github.com/dev-sota/clean-arch-with-caff/utils"
)

// Administrator .
type Administrator struct {
}

// NewAdministrator .
func NewAdministrator() igateway.Administrator {
	return &Administrator{}
}

// Create .
func (a *Administrator) Create(ctx context.Context, administrator *entity.Administrator) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Create(administrator).Error
}

// GetByEmail .
func (a *Administrator) GetByEmail(ctx context.Context, email string) (*entity.Administrator, error) {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return nil, err
	}
	result := &entity.Administrator{}
	err = db.Where("email = ?", email).First(result).Error
	return result, err
}
