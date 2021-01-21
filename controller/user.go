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

type User struct {
	userUsecase iusecase.User
}

func NewUser(userUsecase iusecase.User) icontroller.User {
	return &User{userUsecase: userUsecase}
}

func (u *User) Get(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	ctx := utils.ContextWithDB(gc, infrastructure.GetDB())
	user, err := u.userUsecase.Get(ctx, gc.Param("id"))
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(200, user)
}

func (u *User) List(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	ctx := utils.ContextWithDB(gc, infrastructure.GetDB())
	users, err := u.userUsecase.List(ctx)
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(200, users)
}

func (u *User) Post(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	var user entity.User
	if err := gc.BindJSON(&user); err != nil {
		gc.JSON(500, gin.H{
			"message": "failed to parse JSON",
		})
		return
	}

	err := infrastructure.GetDB().Transaction(func(tx *gorm.DB) error {
		ctx := utils.ContextWithDB(gc, tx)
		return u.userUsecase.Create(ctx, &user)
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

func (u *User) Put(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	var user entity.User
	if err := gc.BindJSON(&user); err != nil {
		gc.JSON(500, gin.H{
			"message": "failed to parse JSON",
		})
		return
	}

	err := infrastructure.GetDB().Transaction(func(tx *gorm.DB) error {
		ctx := utils.ContextWithDB(gc, tx)
		return u.userUsecase.Update(ctx, &user)
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

func (u *User) Delete(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	err := infrastructure.GetDB().Transaction(func(tx *gorm.DB) error {
		ctx := utils.ContextWithDB(gc, tx)
		return u.userUsecase.Delete(ctx, gc.Param("id"))
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
