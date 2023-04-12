package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/willykurniawan01/golang-rest/api/controllers"
)

func InitRoutes(e *echo.Echo){
	// user route
	user := e.Group("users/v1")
	user.GET("",controllers.GetUser)
}