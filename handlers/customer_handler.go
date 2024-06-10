package handlers

import (
	"net/http"
	"strconv"

	"github.com/emr3k0c/money-transfer-system/db"
	"github.com/emr3k0c/money-transfer-system/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateCustomer(c echo.Context) error {
	var customer models.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	createdCustomer, err := db.AddCustomer(customer.Name, customer.Username, customer.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, createdCustomer)
}

func UpdateCustomer(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	customerID := int(claims["customer_id"].(float64))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid customer id"})
	}

	var customer models.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	customer.ID = id

	if customer.ID != customerID {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized action"})
	}

	updatedCustomer, err := db.UpdateCustomer(id, customer.Name, customer.Username, customer.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedCustomer)
}

func ListCustomers(c echo.Context) error {
	customers, err := db.GetCustomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, customers)
}
