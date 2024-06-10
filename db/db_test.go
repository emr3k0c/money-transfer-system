package db_test

import (
	"testing"

	"github.com/emr3k0c/money-transfer-system/db"
	"github.com/stretchr/testify/assert"
)

func TestAddCustomer(t *testing.T) {
	db.InitDB()
	customer, err := db.AddCustomer("Emre Koc", "emrekoc", "password123")
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", customer.Name)
	assert.Equal(t, "johndoe", customer.Username)
}

func TestGetCustomerByUsername(t *testing.T) {
	db.InitDB()
	_, err := db.AddCustomer("Emre Koc", "emrekoc", "password123")
	assert.NoError(t, err)

	customer, err := db.GetCustomerByUsername("emrekoc")
	assert.NoError(t, err)
	assert.Equal(t, "Emre Koc", customer.Name)
}

func TestAddAccount(t *testing.T) {
	db.InitDB()
	customer, err := db.AddCustomer("Emre Koc", "emrekoc", "password123")
	assert.NoError(t, err)

	account, err := db.AddAccount(customer.ID, 1000)
	assert.NoError(t, err)
	assert.Equal(t, customer.ID, account.CustomerID)
	assert.Equal(t, 1000.0, account.Balance)
}

func TestGetAccountsByCustomerID(t *testing.T) {
	db.InitDB()
	customer, err := db.AddCustomer("Emre Koc", "emrekoc", "password123")
	assert.NoError(t, err)

	_, err = db.AddAccount(customer.ID, 1000)
	assert.NoError(t, err)

	accounts, err := db.GetAccountsByCustomerID(customer.ID)
	assert.NoError(t, err)
	assert.Len(t, accounts, 1)
	assert.Equal(t, 1000.0, accounts[0].Balance)
}

func TestTransferMoney(t *testing.T) {
	db.InitDB()
	customer1, err := db.AddCustomer("Emre Koc", "emrekoc", "password123")
	assert.NoError(t, err)

	customer2, err := db.AddCustomer("Berkay Yilmaz", "berkay", "password123")
	assert.NoError(t, err)

	account1, err := db.AddAccount(customer1.ID, 1000)
	assert.NoError(t, err)

	account2, err := db.AddAccount(customer2.ID, 1000)
	assert.NoError(t, err)

	err = db.TransferMoney(account1.ID, account2.ID, 100)
	assert.NoError(t, err)

	updatedAccount1, err := db.GetAccountByID(account1.ID)
	assert.NoError(t, err)
	assert.Equal(t, 900.0, updatedAccount1.Balance)

	updatedAccount2, err := db.GetAccountByID(account2.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1100.0, updatedAccount2.Balance)
}
