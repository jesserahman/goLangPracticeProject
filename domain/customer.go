package domain

import "github.com/jesserahman/goLangPracticeProject/errs"

type Customer struct {
	Id     string `json:"customer_id"`
	Name   string `json:"name"`
	City   string `json:"city"`
	Zip    int    `json:"zip"`
	Status int    `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	ByStatus(string) ([]Customer, *errs.AppError)
}
