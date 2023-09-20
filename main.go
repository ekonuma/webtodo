package main

import (
	"github.com/ekonuma/webtodo/controller"
	"github.com/ekonuma/webtodo/db"
	"github.com/ekonuma/webtodo/repository"
	"github.com/ekonuma/webtodo/router"
	"github.com/ekonuma/webtodo/usecase"
)

func main(){
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	usecase := usecase.NewUserUserCase(userRepository)
	usercontroller := controller.NewUserController(usecase)
	e := router.NewRouter(usercontroller)
	e.Logger.Fatal(e.Start(":8080"))
}