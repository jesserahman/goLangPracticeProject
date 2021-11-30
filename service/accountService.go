package service

import (
	"time"

	"github.com/jesserahman/goLangPracticeProject/domain"
	"github.com/jesserahman/goLangPracticeProject/dto"
	"github.com/jesserahman/goLangPracticeProject/errs"
)

type AccountService interface {
	GetAllAccounts() ([]dto.AccountResponse, *errs.AppError)
	CreateNewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	GetAccountsByCustomerId(string) ([]dto.AccountResponse, *errs.AppError)
	DeleteAccountAndTransactionsByAccountId(accountId string) *errs.AppError
}

type AccountServiceImpl struct {
	repository domain.AccountRepository
}

func (service AccountServiceImpl) GetAllAccounts() ([]dto.AccountResponse, *errs.AppError) {
	accounts, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}
	var accountsDto []dto.AccountResponse
	for _, account := range accounts {
		accountDto := account.ToAccountResponseDto()
		accountsDto = append(accountsDto, *accountDto)
	}
	return accountsDto, nil
}

func (service AccountServiceImpl) GetAccountsByCustomerId(customerId string) ([]dto.AccountResponse, *errs.AppError) {
	accounts, err := service.repository.FindByCustomerId(customerId)
	if err != nil {
		return nil, err
	}
	var accountsDto []dto.AccountResponse
	for _, account := range accounts {
		accountDto := account.ToAccountResponseDto()
		accountsDto = append(accountsDto, *accountDto)
	}
	return accountsDto, nil
}

func (service AccountServiceImpl) DeleteAccountAndTransactionsByAccountId(accountId string) *errs.AppError {
	err := service.repository.DeleteAccountAndTransactions(accountId)
	if err != nil {
		return err
	}
	return nil
}

func (service AccountServiceImpl) CreateNewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	account := domain.Account{
		CustomerId:  request.CustomerId,
		OpeningDate: time.Now().Format("2006-01-01 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
	}
	updatedAccount, err := service.repository.Save(account)
	if err != nil {
		return nil, err
	}

	newAccountResponseDto := updatedAccount.ToNewAccountResponseDto()
	return newAccountResponseDto, nil
}

func NewAccountService(repo domain.AccountRepository) AccountServiceImpl {
	return AccountServiceImpl{repo}
}
