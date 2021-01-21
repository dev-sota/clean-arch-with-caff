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

// Get .
func (a *Administrator) Get(ctx context.Context, id string) (*entity.Administrator, error) {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return nil, err
	}
	result := &entity.Administrator{}
	err = db.Where("id = ?", id).First(result).Error
	return result, err
}

// List .
func (a *Administrator) List(ctx context.Context) ([]*entity.Administrator, error) {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var result []*entity.Administrator
	err = db.Find(&result).Error
	return result, err
}

// Create .
func (a *Administrator) Create(ctx context.Context, administrator *entity.Administrator) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Create(administrator).Error
}

// Update .
func (a *Administrator) Update(ctx context.Context, administrator *entity.Administrator) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Update(administrator).Error
}

// Delete .
func (a *Administrator) Delete(ctx context.Context, id string) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Where("id = ?", id).Delete(&entity.Administrator{}).Error
}
