package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUser(c echo.Context)error{
	return c.JSON(http.StatusOK,map[string]interface{}{
		"name": "John Doe",
        "email": "johndoe@example.com",
	})
}