package routes

import (
	"my_spp/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Trailing Slash for slashing in endpoint
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/users/signup", controllers.SignUpUserController)
	e.POST("/users/login", controllers.LoginUserController)

	return e
}
