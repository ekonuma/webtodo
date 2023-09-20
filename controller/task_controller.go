package controller

import (
	"github.com/ekonuma/webtodo/usecase"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type TaskController struct {
	taskUsecase usecase.ITaskUsecase
}

func NewTaskController(taskUsecase usecase.ITaskUsecase) ITaskController {
	return &TaskController{taskUsecase}
}