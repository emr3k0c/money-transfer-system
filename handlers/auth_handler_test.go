package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emr3k0c/money-transfer-system/db"
	"github.com/emr3k0c/money-transfer-system/handlers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	db.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(`{"username": "johndoe", "password": "password123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
