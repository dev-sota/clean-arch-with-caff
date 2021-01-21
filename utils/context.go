package utils

import (
	"context"

	"github.com/dev-sota/clean-arch-with-caff/consts"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// ContextWithDB .
func ContextWithDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, consts.ContextDB, db)
}

// GetDBFromContext .
func GetDBFromContext(ctx context.Context) (*gorm.DB, error) {
	db, ok := ctx.Value(consts.ContextDB).(*gorm.DB)
	if !ok {
		return nil, errors.New("DB is not found in context")
	}
	if err := db.Error; err != nil {
		return nil, err
	}
	return db, nil
}
