package service

import (
	"time"

	"github.com/jesserahman/goLangPracticeProject/domain"
	"github.com/jesserahman/goLangPracticeProject/dto"
	"github.com/jesserahman/goLangPracticeProject/errs"
)

type AccountService interface {
	GetAllAccounts() ([]domain.Account, *errs.AppError)
	CreateNewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	GetAccountsByCustomerId(string) ([]domain.Account, *errs.AppError)
	GetAccountById(string) (*domain.Account, *errs.AppError)
	DeleteAccountAndTransactionsByAccountId(accountId string) *errs.AppError
	UpdateAccount(dto.UpdateAccountRequest) *errs.AppError
}

type AccountServiceImpl struct {
	repository domain.AccountRepository
}

func (service AccountServiceImpl) GetAllAccounts() ([]domain.Account, *errs.AppError) {
	accounts, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (service AccountServiceImpl) GetAccountsByCustomerId(customerId string) ([]domain.Account, *errs.AppError) {
	accounts, err := service.repository.FindByCustomerId(customerId)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (service AccountServiceImpl) GetAccountById(accountId string) (*domain.Account, *errs.AppError) {
	account, err := service.repository.FindById(accountId)
	if err != nil {
		return nil, err
	}
	return account, nil
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
		Status:      1,
	}
	updatedAccount, err := service.repository.Save(account)
	if err != nil {
		return nil, err
	}

	newAccountResponseDto := updatedAccount.ToNewAccountResponseDto()
	return newAccountResponseDto, nil
}

func (service AccountServiceImpl) UpdateAccount(updateAccountRequest dto.UpdateAccountRequest) *errs.AppError {
	err := updateAccountRequest.Validate()
	if err != nil {
		return err
	}

	account := domain.Account{
		AccountId:   updateAccountRequest.AccountId,
		AccountType: updateAccountRequest.AccountType,
		Status:      updateAccountRequest.Status,
	}

	updateErr := service.repository.Update(account)
	if updateErr != nil {
		return updateErr
	}
	return nil
}

func NewAccountService(repo domain.AccountRepository) AccountServiceImpl {
	return AccountServiceImpl{repo}
}
