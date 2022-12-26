package fcn

import (
	"log"
	"net/http"
	"time"

	"example.com/m/v2/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.Get("username").(string)
	log.Printf("log username: %s\n", username)
	admin := c.Get("admin").(bool)
	log.Printf("log admin: %v\n", admin)

	//create Tonken
	token := jwt.New(jwt.SigningMethodHS256)

	//Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	//genagrated token
	t, err := token.SignedString([]byte("mysecretkey"))
	if err != nil {
		log.Printf("log err token: %v\n", err)
		return err
	}
	// return the token for client
	// return c.JSON(http.StatusOK, echo.Map{
	// 	"token": t,
	// })
	return c.JSON(http.StatusOK, models.LoginResponse{
		Tonken: t,
	})
}
