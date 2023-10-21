package router

import (
	"net/http"
	"os"

	"github.com/ekonuma/webtodo/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func NewRouter(userController controller.IUserController, taskController controller.ITaskController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		CookieSecure: true,
	}))

	e.POST("/signup", userController.SignUp)
	e.POST("/login", userController.LogIn)
	e.POST("/logout", userController.LogOut)
	e.GET("/csrf", userController.CsrfToken)
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
