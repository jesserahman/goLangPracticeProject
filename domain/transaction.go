package domain

import (
	"github.com/jesserahman/goLangPracticeProject/dto"
	"github.com/jesserahman/goLangPracticeProject/errs"
	"github.com/jesserahman/goLangPracticeProject/logger"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type Transaction struct {
	TransactionId   string  `json:"transaction_id" db:"transaction_id"`
	AccountId       string  `json:"account_id" db:"account_id"`
	Amount          float64 `json:"amount" db:"amount"`
	TransactionType string  `json:"transaction_type" db:"transaction_type"`
	TransactionDate string  `json:"transaction_date" db:"transaction_date"`
}

type TransactionRepository interface {
	ExecuteTransaction(Transaction) (*Transaction, *errs.AppError)
	FindByAccountId(string) ([]Transaction, *errs.AppError)
}

func (t Transaction) ToNewTransactionResponseDto() *dto.NewTransactionResponse {
	return &dto.NewTransactionResponse{
		TransactionId: t.TransactionId,
		Balance:       t.Amount,
	}
}

func (t Transaction) ToTransactionResponseDto() *dto.TransactionResponse {
	return &dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func (t Transaction) Validate() *errs.AppError {
	if t.Amount < 0 {
		logger.Error("transaction amount cannot a negative ")
		return errs.NewUnexpectedError("invalid transaction amount")
	}

	if t.TransactionType != WITHDRAWAL && t.TransactionType != DEPOSIT {
		logger.Error("Transaction type must be either 'withdrawal' or 'deposit' ")
		return errs.NewUnexpectedError("invalid transaction type")
	}
	return nil
}
