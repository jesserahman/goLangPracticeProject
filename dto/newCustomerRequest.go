package dto

import (
	"strconv"

	"github.com/jesserahman/goLangPracticeProject/errs"
)

type NewCustomerRequest struct {
	CustomerId  string `json:"customer_id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	City        string `json:"city"`
	ZipCode     int    `json:"zip_code"`
	Status      int    `json:"status"`
}

func (n NewCustomerRequest) Validate() *errs.AppError {
	if len(n.Name) < 1 {
		return errs.NewValidationError("Name must be at least 1 character")
	}

	if n.Status != 1 && n.Status != 0 {
		return errs.NewValidationError("Status must be either 0 or 1")
	}

	if len(strconv.Itoa(n.ZipCode)) != 5 {
		return errs.NewValidationError("Zipcode must be 5 numbers ")
	}

	return nil
}
