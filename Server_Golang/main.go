package main

import (
	"example.com/m/v2/fcn"
	"example.com/m/v2/mdw"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()
	server.Use(middleware.Logger())
	isLoggedIn := middleware.JWT([]byte("mysecretkey"))
	server.POST("/login", fcn.Login, middleware.BasicAuth(mdw.BasicAuth))
	server.GET("/", fcn.Hello, isLoggedIn)
	server.Logger.Fatal(server.Start(":8888"))
}
