package main

import (
	"github.com/willykurniawan01/golang-rest/api/routes"
    "github.com/labstack/echo/v4"
)

func main() {
    // Membuat instance dari Echo
    e := echo.New()
	routes.InitRoutes(e)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}
