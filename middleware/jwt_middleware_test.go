package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/emr3k0c/money-transfer-system/middleware"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestJWTMiddleware(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	os.Setenv("JWT_SECRET", "7bZ!v@Hj9pLw$uQr4NcRt^Fs1KmYxEoT")

	mw := middleware.JWTMiddleware()
	h := mw(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})

	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}
