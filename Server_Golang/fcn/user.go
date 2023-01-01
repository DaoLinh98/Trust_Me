package fcn

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string
	Age  int
}

var listUser = []User{
	{
		Name: "a",
		Age:  1,
	},
	{
		Name: "b",
		Age:  2,
	},
	{
		Name: "c",
		Age:  3,
	},
	{
		Name: "d",
		Age:  4,
	},
}

func GetAllUser(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)

	enc := json.NewEncoder(c.Response())

	for _, user := range listUser {
		if err := enc.Encode(user); err != nil {
			return err
		}

		c.Response().Flush()
		time.Sleep(1 * time.Second)
	}
	return nil
}

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "API get user")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "API UpdateUser user")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "API DeleteUser user")
}
