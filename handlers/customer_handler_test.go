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

func TestCreateCustomer(t *testing.T) {
	db.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/customers", bytes.NewBufferString(`{"name": "John Doe", "username": "johndoe", "password": "password123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.CreateCustomer(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestUpdateCustomer(t *testing.T) {
	db.InitDB()
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/customers", bytes.NewBufferString(`{"name": "John Doe", "username": "johndoe", "password": "password123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, handlers.CreateCustomer(c))
	assert.Equal(t, http.StatusCreated, rec.Code)

	req = httptest.NewRequest(http.MethodPut, "/api/customers/1", bytes.NewBufferString(`{"name": "John Updated", "username": "johnupdated", "password": "newpassword123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, handlers.UpdateCustomer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestListCustomers(t *testing.T) {
	db.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/customers", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.ListCustomers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
