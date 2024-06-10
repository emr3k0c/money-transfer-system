package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emr3k0c/money-transfer-system/db"
	"github.com/emr3k0c/money-transfer-system/handlers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	db.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/accounts", bytes.NewBufferString(`{"customer_id": 1, "balance": 1000}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": float64(1),
	})
	c.Set("user", token)

	if assert.NoError(t, handlers.CreateAccount(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestUpdateAccount(t *testing.T) {
	db.InitDB()
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/accounts", bytes.NewBufferString(`{"customer_id": 1, "balance": 1000}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": float64(1),
	})
	c.Set("user", token)

	assert.NoError(t, handlers.CreateAccount(c))
	assert.Equal(t, http.StatusCreated, rec.Code)

	req = httptest.NewRequest(http.MethodPut, "/api/accounts/1", bytes.NewBufferString(`{"balance": 1500}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("user", token)

	if assert.NoError(t, handlers.UpdateAccount(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestTransferMoney(t *testing.T) {
	db.InitDB()
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/accounts", bytes.NewBufferString(`{"customer_id": 1, "balance": 1000}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": float64(1),
	})
	c.Set("user", token)
	c.Set("user", token)

	assert.NoError(t, handlers.CreateAccount(c))
	assert.Equal(t, http.StatusCreated, rec.Code)

	req = httptest.NewRequest(http.MethodPost, "/api/accounts", bytes.NewBufferString(`{"customer_id": 1, "balance": 500}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.Set("user", token)

	assert.NoError(t, handlers.CreateAccount(c))
	assert.Equal(t, http.StatusCreated, rec.Code)

	req = httptest.NewRequest(http.MethodPost, "/api/accounts/transfer", bytes.NewBufferString(`{"from_account_id": 1, "to_account_id": 2, "amount": 100}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.Set("user", token)

	if assert.NoError(t, handlers.TransferMoney(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestListAccounts(t *testing.T) {
	db.InitDB()
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/accounts", bytes.NewBufferString(`{"customer_id": 1, "balance": 1000}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	user := map[string]interface{}{
		"customer_id": 1,
	}
	c.Set("user", user)

	assert.NoError(t, handlers.CreateAccount(c))
	assert.Equal(t, http.StatusCreated, rec.Code)

	req = httptest.NewRequest(http.MethodGet, "/api/accounts", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.Set("user", user)

	if assert.NoError(t, handlers.ListAccounts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
