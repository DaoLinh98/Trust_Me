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
	isAdmin := mdw.IsAdmin
	server.POST("/login", fcn.Login, middleware.BasicAuth(mdw.BasicAuth))
	server.GET("/", fcn.Hello, isLoggedIn)
	server.GET("/admin", fcn.Hello, isLoggedIn, isAdmin)

	groupv2 := server.Group("/v2")
	groupv2.GET("/hello", fcn.Hello2)

	groupUser := server.Group("/api/user")
	groupUser.GET("/GetUser", fcn.GetUser)
	groupUser.GET("/GetAllUser", fcn.GetAllUser)
	groupUser.GET("/UpdateUser", fcn.UpdateUser, isAdmin)
	groupUser.GET("/DeleteUser", fcn.DeleteUser, isAdmin)

	server.Logger.Fatal(server.Start(":8888"))
}
