package domain

import (
	"fmt"
	"github.com/jesserahman/goLangPracticeProject/errs"
	"github.com/jesserahman/goLangPracticeProject/logger"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type TransactionRepositoryDb struct {
	dbClient *sqlx.DB
}

func (t TransactionRepositoryDb)ExecuteTransaction(transaction Transaction) (*Transaction, *errs.AppError){
	// update bank account
	accountUpdate := fmt.Sprintf("Update banking.accounts Set amount = %f WHERE account_id = '%s'", transaction.Amount, transaction.AccountId)
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
