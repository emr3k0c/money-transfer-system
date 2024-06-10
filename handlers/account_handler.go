package handlers

import (
	"net/http"
	"strconv"

	"github.com/emr3k0c/money-transfer-system/db"
	"github.com/emr3k0c/money-transfer-system/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateAccount(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	customerID := int(claims["customer_id"].(float64))

	var account models.Account
	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if account.CustomerID != customerID {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized action"})
	}

	createdAccount, err := db.AddAccount(account.CustomerID, account.Balance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, createdAccount)
}

func UpdateAccount(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	customerID := int(claims["customer_id"].(float64))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid account id"})
	}

	var account models.Account
	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if account.CustomerID != customerID {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized action"})
	}

	updatedAccount, err := db.UpdateAccount(id, account.Balance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedAccount)
}

func TransferMoney(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	customerID := int(claims["customer_id"].(float64))

	type transferRequest struct {
		FromAccountID int     `json:"from_account_id"`
		ToAccountID   int     `json:"to_account_id"`
		Amount        float64 `json:"amount"`
	}
	var req transferRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	fromAccount, err := db.GetAccountByID(req.FromAccountID)
	if err != nil || fromAccount.CustomerID != customerID {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized action"})
	}

	_, err = db.GetAccountByID(req.ToAccountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	err = db.TransferMoney(req.FromAccountID, req.ToAccountID, req.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "transfer successful"})
}

func ListAccounts(c echo.Context) error {
	user := c.Get("user").(jwt.MapClaims)
	customerID := int(user["customer_id"].(float64))

	accounts, err := db.GetAccountsByCustomerID(customerID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, accounts)
}
