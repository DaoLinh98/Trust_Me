package fcn

import (
	"net/http"

	"example.com/m/v2/models"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	x := &models.X{
		Text: "hello"}

	return c.JSON(http.StatusOK, x)
}
