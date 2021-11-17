package service

import (
	"github.com/jesserahman/goLangPracticeProject/domain"
	"github.com/jesserahman/goLangPracticeProject/dto"
	"github.com/jesserahman/goLangPracticeProject/errs"
	"time"
)

type TransactionService interface {
	CreateNewTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type TransactionServiceImpl struct {
	repository domain.TransactionRepository
}

func (service TransactionServiceImpl) CreateNewTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	transaction := domain.Transaction{
		AccountId:       request.AccountId,
		Amount:          request.Amount,
		TransactionType: request.TransactionType,
		TransactionDate: time.Now().Format("2006-01-01 15:04:05"),
	}
	updatedTransaction, err := service.repository.ExecuteTransaction(transaction)
	if err != nil {
		return nil, err
	}

	transactionResponseDto := updatedTransaction.ToTransactionResponseDto()

	return transactionResponseDto, nil
}

func NewTransactionService(repo domain.TransactionRepository) TransactionServiceImpl {
	return TransactionServiceImpl{repo}
}