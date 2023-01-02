package fcn

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

func init() {
	orm.RegisterModel(new(User))

}

type User struct {
	Id    int64  `orm:"auto json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
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

func AddUser(c echo.Context) error {
	user := &User{}
	if err := c.Bind(user); err != nil {
		glog.Error("bind user error: %v", err)
		return err
	}

	o := orm.NewOrm()
	id, err := o.Insert(user)
	if err != nil {
		glog.Error("inser user error : %v", err)
		return err
	}
	glog.Infof("insert at row %d", id)
	return c.JSON(http.StatusOK, user)

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
