package middleware

import (
	"log"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(secret),
		TokenLookup: "header:Authorization",
		ContextKey:  "user",
		SuccessHandler: func(c echo.Context) {
		},
		ErrorHandler: func(c echo.Context, err error) error {
			log.Println("Token verification failed:", err)
			return err
		},
	})
}
