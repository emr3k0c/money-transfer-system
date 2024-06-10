package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/emr3k0c/money-transfer-system/models"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	_, err = DB.Exec(`CREATE TABLE customers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatalf("Failed to create customers table: %v", err)
	}

	_, err = DB.Exec(`CREATE TABLE accounts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_id INTEGER,
		balance REAL,
		FOREIGN KEY (customer_id) REFERENCES customers(id)
	)`)
	if err != nil {
		log.Fatalf("Failed to create accounts table: %v", err)
	}
}

func GetCustomers() ([]models.Customer, error) {
	rows, err := DB.Query("SELECT id, name, username FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Username); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func AddCustomer(name, username, password string) (models.Customer, error) {
	result, err := DB.Exec("INSERT INTO customers (name, username, password) VALUES (?, ?, ?)", name, username, password)
	if err != nil {
		return models.Customer{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return models.Customer{}, err
	}

	return models.Customer{ID: int(id), Name: name, Username: username, Password: password}, nil
}

func UpdateCustomer(id int, name, username, password string) (models.Customer, error) {
	result, err := DB.Exec("UPDATE customers SET name = ?, username = ?, password = ? WHERE id = ?", name, username, password, id)
	if err != nil {
		return models.Customer{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Customer{}, err
	}
	if rowsAffected == 0 {
		return models.Customer{}, errors.New("customer not found")
	}
	return models.Customer{ID: id, Name: name, Username: username, Password: password}, nil
}

func GetCustomerByUsername(username string) (models.Customer, error) {
	var customer models.Customer
	err := DB.QueryRow("SELECT id, name, username, password FROM customers WHERE username = ?", username).Scan(&customer.ID, &customer.Name, &customer.Username, &customer.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, errors.New("customer not found")
		}
		return customer, err
	}
	return customer, nil
}

func AddAccount(customerID int, balance float64) (models.Account, error) {
	result, err := DB.Exec("INSERT INTO accounts (customer_id, balance) VALUES (?, ?)", customerID, balance)
	if err != nil {
		return models.Account{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return models.Account{}, err
	}
	return models.Account{ID: int(id), CustomerID: customerID, Balance: balance}, nil
}

func UpdateAccount(id int, balance float64) (models.Account, error) {
	result, err := DB.Exec("UPDATE accounts SET balance = ? WHERE id = ?", balance, id)
	if err != nil {
		return models.Account{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Account{}, err
	}
	if rowsAffected == 0 {
		return models.Account{}, errors.New("account not found")
	}
	return models.Account{ID: id, Balance: balance}, nil
}

func GetAccountsByCustomerID(customerID int) ([]models.Account, error) {
	rows, err := DB.Query("SELECT id, customer_id, balance FROM accounts WHERE customer_id = ?", customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []models.Account
	for rows.Next() {
		var account models.Account
		if err := rows.Scan(&account.ID, &account.CustomerID, &account.Balance); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func GetAccountByID(accountID int) (models.Account, error) {
	var account models.Account
	row := DB.QueryRow("SELECT id, customer_id, balance FROM accounts WHERE id = ?", accountID)
	err := row.Scan(&account.ID, &account.CustomerID, &account.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return account, errors.New("account not found")
		}
		return account, err
	}
	return account, nil
}

func TransferMoney(fromAccountID, toAccountID int, amount float64) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var fromBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", fromAccountID).Scan(&fromBalance)
	if err != nil {
		return err
	}
	if fromBalance < amount {
		return errors.New("insufficient funds")
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromAccountID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toAccountID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
