package router

import (
	"github.com/ekonuma/webtodo/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(userController controller.IUserController) *echo.Echo{
	e := echo.New()
	e.POST("/signup", userController.SignUp)
	e.POST("/login", userController.LogIn)
	e.POST("/logout", userController.LogOut)
	return e
}