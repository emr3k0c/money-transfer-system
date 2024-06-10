package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/emr3k0c/money-transfer-system/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	customer, err := db.GetCustomerByUsername(req.Username)
	if err != nil || customer.Password != req.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["customer_id"] = customer.ID
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": t})

}
