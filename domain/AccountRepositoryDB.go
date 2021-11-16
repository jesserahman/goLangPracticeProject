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
	// query the DB, and store the result in ${customers}
	err := a.dbClient.Select(&accounts, accountsQuery)
	if err != nil {
		logger.Error("Error querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return accounts, nil
}

func (a AccountRepositoryDb) FindByCustomerId(customerId string) ([]Account, *errs.AppError) {
	accounts := make([]Account, 0)
	accountsQuery := fmt.Sprintf("select * from banking.accounts where customer_id = %s", customerId)
	// query the DB, and store the result in ${customers}
	err := a.dbClient.Select(&accounts, accountsQuery)
	if err != nil {
		logger.Error("Error querying customers table " + err.Error())
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

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error getting last inserted ID" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	account.AccountId = strconv.FormatInt(id, 10)
	return &account, nil
}

func NewAccountRepositoryDbConnection(dbClient *sqlx.DB) AccountRepository {
	return AccountRepositoryDb{dbClient}
}
