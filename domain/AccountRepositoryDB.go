package domain

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jesserahman/goLangPracticeProject/errs"
	"github.com/jesserahman/goLangPracticeProject/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	dbClient *sqlx.DB
}

func (a AccountRepositoryDb) FindAll() ([]Account, *errs.AppError) {
	accounts := make([]Account, 0)
	accountsQuery := "select * from accounts"
	// query the DB, and store the result in ${accounts}
	err := a.dbClient.Select(&accounts, accountsQuery)
	if err != nil {
		logger.Error("Error querying Accounts table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return accounts, nil
}

func (a AccountRepositoryDb) FindById(accountId string) (*Account, *errs.AppError) {
	accountsQuery := fmt.Sprintf("select * from accounts where account_id = %s", accountId)

	var account Account
	// query the DB, and store the result in var account
	err := a.dbClient.Get(&account, accountsQuery)
	if err != nil {
		logger.Error("Error querying Accounts table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return &account, nil
}

func (a AccountRepositoryDb) FindByCustomerId(customerId string) ([]Account, *errs.AppError) {
	accounts := make([]Account, 0)
	accountsQuery := fmt.Sprintf("select * from banking.accounts where customer_id = %s", customerId)

	// query the DB, and store the result in var accounts
	err := a.dbClient.Select(&accounts, accountsQuery)
	if err != nil {
		logger.Error("Error querying Accounts table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return accounts, nil
}

func (a AccountRepositoryDb) Save(account Account) (*Account, *errs.AppError) {
	accountsInsert := "INSERT into accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	result, err := a.dbClient.Exec(accountsInsert, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	log.Println("result: ", result)
	if err != nil {
		logger.Error("Error inserting into Accounts table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	id, resultErr := result.LastInsertId()
	if resultErr != nil {
		logger.Error("Error getting last inserted ID" + resultErr.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	account.AccountId = strconv.FormatInt(id, 10)
	return &account, nil
}

func (a AccountRepositoryDb) Update(account Account) *errs.AppError {
	accountsUpdate := "Update accounts Set account_type=?, status=? where account_id = ?"
	_, err := a.dbClient.Exec(accountsUpdate, account.AccountType, account.Status, account.AccountId)

	if err != nil {
		logger.Error("Error Updating Accounts table " + err.Error())
		return errs.NewUnexpectedError("unexpected database error")
	}

	return nil
}

func (a AccountRepositoryDb) DeleteAccountAndTransactions(accountId string) *errs.AppError {
	// starting database transaction block
	tx, err := a.dbClient.Begin()
	if err != nil {
		logger.Error("Error starting the db transaction block " + err.Error())
		return errs.NewUnexpectedError("unexpected database error")
	}

	// delete all transactions for that account
	transactionsDelete := fmt.Sprintf("DELETE FROM transactions WHERE account_id = %s", accountId)
	_, transacationsDeleteErr := tx.Exec(transactionsDelete)
	if transacationsDeleteErr != nil {
		tx.Rollback()
		logger.Error("Error deleting from Transactions table " + transacationsDeleteErr.Error())
		return errs.NewUnexpectedError("unexpected database error")
	}

	// delete account from accounts table
	accountDelete := fmt.Sprintf("DELETE FROM accounts WHERE account_id = %s", accountId)
	_, accountDeleteErr := tx.Exec(accountDelete)
	if accountDeleteErr != nil {
		tx.Rollback()
		logger.Error("Error deleting from Accounts table " + accountDeleteErr.Error())
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

func NewAccountRepositoryDbConnection(dbClient *sqlx.DB) AccountRepository {
	return AccountRepositoryDb{dbClient}
}
