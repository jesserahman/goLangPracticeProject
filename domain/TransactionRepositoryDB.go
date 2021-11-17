package domain

import (
	"fmt"
	"github.com/jesserahman/goLangPracticeProject/errs"
	"github.com/jesserahman/goLangPracticeProject/logger"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
	"strings"
)

type TransactionRepositoryDb struct {
	dbClient *sqlx.DB
}

func (t TransactionRepositoryDb)ExecuteTransaction(transaction Transaction) (*Transaction, *errs.AppError){
	// get current bank account details
	accountQuery := fmt.Sprintf("select * from banking.accounts where account_id = %s", transaction.AccountId)
	account := Account{AccountId: transaction.AccountId}
	err := t.dbClient.QueryRowx(accountQuery).StructScan(&account)
	if err != nil {
		logger.Error("Error querying accounts table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// update bank account
	var newBalance float64
	if strings.ToLower(transaction.TransactionType) == "withdraw" {
		newBalance = account.Amount - transaction.Amount
	} else if strings.ToLower(transaction.TransactionType) == "deposit" {
		newBalance = account.Amount + transaction.Amount
	} else {
		return nil, errs.NewUnexpectedError("invalid transaction type")
	}
	accountUpdate := fmt.Sprintf("Update banking.accounts Set amount = %f WHERE account_id = '%s'", newBalance, transaction.AccountId)
	result, err := t.dbClient.Exec(accountUpdate, )
	log.Println("result: ", result)
	if err != nil {
		logger.Error("Error updating Accounts table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error getting last inserted ID" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	transaction.AccountId = strconv.FormatInt(id, 10)

	return &transaction, nil
}

func NewTransactionRepositoryDbConnection(dbClient *sqlx.DB) TransactionRepository {
	return TransactionRepositoryDb{dbClient}
}
