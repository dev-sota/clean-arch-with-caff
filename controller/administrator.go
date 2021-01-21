package controller

import (
	"context"

	"github.com/dev-sota/clean-arch-with-caff/entity"
	"github.com/dev-sota/clean-arch-with-caff/icontroller"
	"github.com/dev-sota/clean-arch-with-caff/infrastructure"
	"github.com/dev-sota/clean-arch-with-caff/iusecase"
	"github.com/dev-sota/clean-arch-with-caff/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Administrator struct {
	administratorUsecase iusecase.Administrator
}

func NewAdministrator(administratorUsecase iusecase.Administrator) icontroller.Administrator {
	return &Administrator{administratorUsecase: administratorUsecase}
}

func (a *Administrator) Get(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	ctx := utils.ContextWithDB(gc, infrastructure.GetDB())
	administrator, err := a.administratorUsecase.Get(ctx, gc.Param("id"))
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(200, administrator)
}

func (a *Administrator) List(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	ctx := utils.ContextWithDB(gc, infrastructure.GetDB())
	administrators, err := a.administratorUsecase.List(ctx)
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(200, administrators)
}

func (a *Administrator) Post(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	var administrator entity.Administrator
	if err := gc.BindJSON(&administrator); err != nil {
		gc.JSON(500, gin.H{
			"message": "failed to parse JSON",
		})
		return
	}

	err := infrastructure.GetDB().Transaction(func(tx *gorm.DB) error {
		ctx := utils.ContextWithDB(gc, tx)
		return a.administratorUsecase.Create(ctx, &administrator)
	})
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(201, gin.H{
		"message": "created",
	})
}

func (a *Administrator) Put(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	var administrator entity.Administrator
	if err := gc.BindJSON(&administrator); err != nil {
		gc.JSON(500, gin.H{
			"message": "failed to parse JSON",
		})
		return
	}

	err := infrastructure.GetDB().Transaction(func(tx *gorm.DB) error {
		ctx := utils.ContextWithDB(gc, tx)
		return a.administratorUsecase.Update(ctx, &administrator)
	})
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(200, gin.H{
		"message": "updated",
	})
}

func (a *Administrator) Delete(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	err := infrastructure.GetDB().Transaction(func(tx *gorm.DB) error {
		ctx := utils.ContextWithDB(gc, tx)
		return a.administratorUsecase.Delete(ctx, gc.Param("id"))
	})
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(200, gin.H{
		"message": "deleted",
	})
}
