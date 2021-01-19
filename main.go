package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/dev-sota/clean-arch-with-caff/controller"
	"github.com/dev-sota/clean-arch-with-caff/gateway"
	"github.com/dev-sota/clean-arch-with-caff/infrastructure"
	"github.com/dev-sota/clean-arch-with-caff/model"
	"github.com/dev-sota/clean-arch-with-caff/usecase"
	"os"
)

func main() {
	mysqlConfig := model.NewMysql("root", "", "localhost", 3306, "hoge")

	if err := infrastructure.InitDB(mysqlConfig.String()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := gin.Default()

	userGateway := gateway.NewUser()
	userUsecase := usecase.NewUser(userGateway)
	userController := controller.NewUser(userUsecase)
	userRooter := r.Group("users")
	userRooter.GET("", toGin(userController.List))
	userRooter.GET(":id", toGin(userController.Get))
	userRooter.POST("", toGin(userController.Post))
	userRooter.PUT(":id", toGin(userController.Put))
	userRooter.DELETE(":id", toGin(userController.Delete))

	r.Run()
}

func toGin(f func(context.Context)) func(*gin.Context) {
	return func(gc *gin.Context) {
		f(gc)
	}
}