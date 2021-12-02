package domain

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/jesserahman/goLangPracticeProject/logger"
	"github.com/jmoiron/sqlx"

	"github.com/jesserahman/goLangPracticeProject/errs"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	dbClient *sqlx.DB
}

func (c CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var customer Customer
	customersQuery := fmt.Sprintf("select * from customers where customer_id = '%s'", id)
	err := c.dbClient.Get(&customer, customersQuery)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error searching DB for customer_id " + err.Error())
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error searching DB for customer_id " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &customer, nil
}

func (c CustomerRepositoryDb) ByStatus(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	customerStatus, _ := strconv.Atoi(status)
	customersQuery := fmt.Sprintf("select * from customers where status = %d", customerStatus)

	// query the DB, and store the result in ${customers}
	err := c.dbClient.Select(&customers, customersQuery)
	if err != nil {
		logger.Error("Error searching by status in the customers table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	customersQuery := "select * from customers"
	// query the DB, and store the result in ${customers}
	err := c.dbClient.Select(&customers, customersQuery)
	if err != nil {
		logger.Error("Error querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (c CustomerRepositoryDb) Save(customer Customer) (*Customer, *errs.AppError) {
	customerInsert := "INSERT into customers (name, date_of_birth, city, zipcode, status) VALUES (?, ?, ?, ?, ?)"
	result, err := c.dbClient.Exec(customerInsert, customer.Name, customer.DateOfBirth, customer.City, customer.Zip, customer.Status)
	if err != nil {
		logger.Error("Error inserting into Customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	id, resultErr := result.LastInsertId()
	if resultErr != nil {
		logger.Error("Error getting Customer ID")
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	customer.Id = strconv.FormatInt(id, 10)

	return &customer, nil
}

func (c CustomerRepositoryDb) Update(customer Customer) (*Customer, *errs.AppError) {
	customerUpdate := "Update customers Set name=?, date_of_birth=?, city=?, zipcode=?, status=? where customer_id = ?"
	_, err := c.dbClient.Exec(customerUpdate, customer.Name, customer.DateOfBirth, customer.City, customer.Zip, customer.Status, customer.Id)
	if err != nil {
		logger.Error("Error updating Customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return &customer, nil
}

func (c CustomerRepositoryDb) Delete(customerId string) *errs.AppError {
	// Get a list of all customer accounts
	accounts := make([]Account, 0)
	accountsQuery := fmt.Sprintf("select * from banking.accounts where customer_id = %s", customerId)
	log.Printf("Accounts: %v\n", accounts)

	// query the DB, and store the result in var accounts
	err := c.dbClient.Select(&accounts, accountsQuery)
	if err != nil {
		logger.Error("Error querying Accounts table " + err.Error())
		return errs.NewUnexpectedError("unexpected database error")
	}

	// starting database transaction block
	tx, err := c.dbClient.Begin()
	if err != nil {
		logger.Error("Error starting the db transaction block " + err.Error())
		return errs.NewUnexpectedError("unexpected database error")
	}

	// loop through all accounts, and delete all transactions and the accounts
	for _, account := range accounts {
		//	for each account delete all transactions

		// delete all transactions for that account
		transactionsDelete := fmt.Sprintf("DELETE FROM transactions WHERE account_id = %s", account.AccountId)
		_, transacationsDeleteErr := tx.Exec(transactionsDelete)
		if transacationsDeleteErr != nil {
			tx.Rollback()
			logger.Error("Error deleting from Transactions table " + transacationsDeleteErr.Error())
			return errs.NewUnexpectedError("unexpected database error")
		}

		// delete account from accounts table
		accountDelete := fmt.Sprintf("DELETE FROM accounts WHERE account_id = %s", account.AccountId)
		_, accountDeleteErr := tx.Exec(accountDelete)
		if accountDeleteErr != nil {
			tx.Rollback()
			logger.Error("Error deleting from Accounts table " + accountDeleteErr.Error())
			return errs.NewUnexpectedError("unexpected database error")
		}

	}

	// delete customer from customers table
	customerDelete := fmt.Sprintf("DELETE FROM customers WHERE customer_id = %s", customerId)
	_, customerDeleteErr := tx.Exec(customerDelete)
	if customerDeleteErr != nil {
		tx.Rollback()
		logger.Error("Error deleting from Customers table " + customerDeleteErr.Error())
		return errs.NewUnexpectedError("unexpected database error")
	}

	// if there are no errors then commit the change
	commitErr := tx.Commit()
	if commitErr != nil {
		tx.Rollback()
		logger.Error("Error committing changes" + commitErr.Error())
		return errs.NewUnexpectedError("unexpected database error")
	}
	return nil
}

func NewCustomerRepositoryDbConnection(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
