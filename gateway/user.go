package gateway

import (
	"context"
	"github.com/dev-sota/clean-arch-with-caff/igateway"
	"github.com/dev-sota/clean-arch-with-caff/entity"
	"github.com/dev-sota/clean-arch-with-caff/utils"
)

// User .
type User struct {
}

// NewUser .
func NewUser() igateway.User {
	return &User{}
}

// Get .
func (u *User) Get(ctx context.Context, id string) (*entity.User, error) {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return nil, err
	}
	result := &entity.User{}
	err = db.Where("id = ?", id).First(result).Error
	return result, err
}

// List .
func (u *User) List(ctx context.Context) ([]*entity.User, error) {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var result []*entity.User
	err = db.Find(&result).Error
	return result, err
}

// Create .
func (u *User) Create(ctx context.Context, user *entity.User) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Create(user).Error
}

// Update .
func (u *User) Update(ctx context.Context, user *entity.User) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Update(user).Error
}

// Delete .
func (u *User) Delete(ctx context.Context, id string) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Where("id = ?", id).Delete(&entity.User{}).Error
}
