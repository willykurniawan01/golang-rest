package routes

import (
	"github.com/labstack/echo/v4"
	"fmt"
	"github.com/willykurniawan01/golang-rest/api/controllers"
) 

func InitUserRoutes(e *echo.Echo, c *controllers.Controller) {
	user := e.Group("/api/v1/users")
	user.GET("",c.Getuser)
}