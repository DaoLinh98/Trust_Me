package main

import (
	"log"

	"example.com/m/v2/fcn"
	"example.com/m/v2/mdw"
	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	// register model
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	err := orm.RegisterDataBase("default", "mysql", "root:123456@/echo_test?charset=utf8", 30)
	if err != nil {
		log.Printf("Failed to register database %v", err)
	}
	//alias
	name := "default"

	//drop table and re-create
	force := false

	//prnit log mysql
	verbose := true

	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		log.Print(err)
	}
	// create table
	//orm.RunSyncdb("default", false, true)
}

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
	groupUser.PUT("/add", fcn.AddUser)
	groupUser.GET("/GetUser", fcn.GetUser)
	groupUser.GET("/GetAllUser", fcn.GetAllUser)
	groupUser.GET("/UpdateUser", fcn.UpdateUser, isAdmin)
	groupUser.GET("/DeleteUser", fcn.DeleteUser, isAdmin)

	server.Logger.Fatal(server.Start(":8888"))
}
