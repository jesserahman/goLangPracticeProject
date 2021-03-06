package domain

import (
	"github.com/jesserahman/goLangPracticeProject/dto"
	"github.com/jesserahman/goLangPracticeProject/errs"
)

type Customer struct {
	Id          string `json:"customer_id" db:"customer_id"`
	Name        string `json:"name" db:"name"`
	City        string `json:"city" db:"city"`
	Zip         int    `json:"zip_code" db:"zipcode"`
	Status      int    `json:"status" db:"status"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	ByStatus(string) ([]Customer, *errs.AppError)
	Save(Customer) (*Customer, *errs.AppError)
	Update(Customer) (*Customer, *errs.AppError)
	Delete(customerId string) *errs.AppError
}

func (c Customer) getStatus() string {
	status := "active"
	if c.Status == 0 {
		status = "inactive"
	}
	return status
}

func (c Customer) ToCustomerResponseDto() *dto.CustomerResponse {
	return &dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zip:         c.Zip,
		Status:      c.getStatus(),
		DateOfBirth: c.DateOfBirth,
	}
}

func (c Customer) ToNewCustomerResponseDto() *dto.NewCustomerResponse {
	return &dto.NewCustomerResponse{CustomerId: c.Id}
}
