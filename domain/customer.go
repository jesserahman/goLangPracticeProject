package domain

import "github.com/jesserahman/goLangPracticeProject/errs"

type Customer struct {
	Id     string `json:"customer_id" db:"customer_id"`
	Name   string `json:"name" db:"name"`
	City   string `json:"city" db:"city"`
	Zip    int    `json:"zip" db:"zipcode"`
	Status int    `json:"status" db:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	ByStatus(string) ([]Customer, *errs.AppError)
}
