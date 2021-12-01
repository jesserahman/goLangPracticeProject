package dto

import (
	"strings"

	"github.com/jesserahman/goLangPracticeProject/errs"
)

type UpdateAccountRequest struct {
	AccountId   string `json:"account_id"`
	AccountType string `json:"account_type"`
	Status      int    `json:"status"`
}

func (u UpdateAccountRequest) Validate() *errs.AppError {
	if u.Status != 0 && u.Status != 1 {
		return errs.NewValidationError("Status must be 0 or 1")
	}
	if strings.ToLower(u.AccountType) != "savings" && strings.ToLower(u.AccountType) != "checking" {
		return errs.NewValidationError("Account type must be either checking or savings")
	}
	return nil
}
