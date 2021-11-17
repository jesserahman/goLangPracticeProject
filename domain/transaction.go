package domain

import (
	"github.com/jesserahman/goLangPracticeProject/dto"
	"github.com/jesserahman/goLangPracticeProject/errs"
)

type Transaction struct {
	TransactionId  string  `json:"transaction_id" db:"transaction_id"`
	AccountId   string  `json:"account_id" db:"account_id"`
	Amount      float64 `json:"amount" db:"amount"`
	TransactionType string `json:"transaction_type" db:"transaction_type"`
	TransactionDate string `json:"transaction_date" db:"transaction_date"`
}

type TransactionRepository interface {
	ExecuteTransaction(transaction Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction)ToTransactionResponseDto() *dto.NewTransactionResponse{
	return &dto.NewTransactionResponse{
		TransactionId: t.TransactionId,
		Balance:       t.Amount,
	}
}
