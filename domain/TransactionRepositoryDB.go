package domain

import (
	"fmt"
	"github.com/jesserahman/goLangPracticeProject/errs"
	"github.com/jesserahman/goLangPracticeProject/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

type TransactionRepositoryDb struct {
	dbClient *sqlx.DB
}

func (t TransactionRepositoryDb)ExecuteTransaction(transaction Transaction) (*Transaction, *errs.AppError){
	// check if amount is negative
	if transaction.Amount < 0 {
		logger.Error("Negative transaction amount")
		return nil, errs.NewUnexpectedError("transaction account cannot be negative")
	}

	// get current bank account details
	accountQuery := fmt.Sprintf("select * from banking.accounts where account_id = %s", transaction.AccountId)
	account := Account{AccountId: transaction.AccountId}
	err := t.dbClient.QueryRowx(accountQuery).StructScan(&account)
	if err != nil {
		logger.Error("Error querying accounts table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// determine new bank account balance
	var newBalance float64
	if strings.ToLower(transaction.TransactionType) == "withdraw" {
		//verify withdraw amount is not greater than total amount in the account
		if transaction.Amount > account.Amount{
			logger.Error("Withdraw amount exceeds total account balance")
			return nil, errs.NewUnexpectedError("withdraw amount exceeds total account balance")
		}
		newBalance = account.Amount - transaction.Amount
	} else if strings.ToLower(transaction.TransactionType) == "deposit" {
		newBalance = account.Amount + transaction.Amount
	} else {
		return nil, errs.NewUnexpectedError("invalid transaction type")
	}

	// Update bank account with new balance
	accountUpdateQuery := fmt.Sprintf("Update banking.accounts Set amount = %f WHERE account_id = '%s'", newBalance, transaction.AccountId)
	_, err = t.dbClient.Exec(accountUpdateQuery)
	if err != nil {
		logger.Error("Error updating Accounts table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// insert transaction into transactions table
	transactionsInsert := "INSERT into banking.transactions (account_id, amount, transaction_type) VALUES (?, ?, ?)"
	result, dbErr := t.dbClient.Exec(transactionsInsert, transaction.AccountId, transaction.Amount, transaction.TransactionType)
	if dbErr != nil {
		logger.Error("Error inserting into Transactions table " + dbErr.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// get transaction ID
	id, resultErr := result.LastInsertId()
	if resultErr != nil {
		logger.Error("Error getting last inserted ID" + resultErr.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	transaction.TransactionId = strconv.FormatInt(id, 10)

	// set transaction amount as new account balance to get returned to the user
	transaction.Amount = newBalance

	return &transaction, nil
}

func (t TransactionRepositoryDb)FindByAccountId(accountId string) ([]Transaction, *errs.AppError){
	transactions := make([]Transaction, 0)
	transactionsQuery := fmt.Sprintf("select * from banking.transactions where account_id = %s", accountId)

	// query the DB, and store the result in transactions
	err := t.dbClient.Select(&transactions, transactionsQuery)
	if err != nil {
		logger.Error("Error querying transactions table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return transactions, nil
}


func NewTransactionRepositoryDbConnection(dbClient *sqlx.DB) TransactionRepository {
	return TransactionRepositoryDb{dbClient}
}
