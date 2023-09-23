package main

import (
	"log"

	"github.com/ekonuma/webtodo/controller"
	"github.com/ekonuma/webtodo/db"
	"github.com/ekonuma/webtodo/repository"
	"github.com/ekonuma/webtodo/router"
	"github.com/ekonuma/webtodo/usecase"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUserCase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
