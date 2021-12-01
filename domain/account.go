package domain

import (
	"github.com/jesserahman/goLangPracticeProject/dto"
	"github.com/jesserahman/goLangPracticeProject/errs"
)

type Account struct {
	AccountId   string  `json:"account_id" db:"account_id"`
	CustomerId  string  `json:"customer_id" db:"customer_id"`
	OpeningDate string  `json:"opening_date" db:"opening_date"`
	AccountType string  `json:"account_type" db:"account_type"`
	Amount      float64 `json:"amount" db:"amount"`
	Status      int     `json:"status" db:"status"`
}

type AccountRepository interface {
	FindAll() ([]Account, *errs.AppError)
	Save(Account) (*Account, *errs.AppError)
	FindById(accountId string) (*Account, *errs.AppError)
	FindByCustomerId(string) ([]Account, *errs.AppError)
	DeleteAccountAndTransactions(string) *errs.AppError
	Update(account Account) *errs.AppError
}

func (a Account) ToAccountResponseDto() *dto.AccountResponse {
	return &dto.AccountResponse{
		AccountId:   a.AccountId,
		CustomerId:  a.CustomerId,
		OpeningDate: a.OpeningDate,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      a.Status,
	}
}
func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
