package router

import (
	"os"

	"github.com/ekonuma/webtodo/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(userController controller.IUserController, taskController controller.ITaskController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", userController.SignUp)
	e.POST("/login", userController.LogIn)
	e.POST("/logout", userController.LogOut)

	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", taskController.GetAllTasks)
	t.GET("/:taskId", taskController.GetTaskById)
	t.POST("", taskController.CreateTask)
	t.PUT("/:taskId", taskController.UpdateTask)
	t.DELETE("/:taskId", taskController.DeleteTask)
	return e
}
