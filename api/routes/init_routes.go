package routes 

import (
	"github.com/labstack/echo/v4"
	"github.com/willykurniawan01/golang-rest/controllers"
)

func InitRoutes(e *echo.Echo,c *controllers.controllers){
	InitUserRoutes(e,c)
}