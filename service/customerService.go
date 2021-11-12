package service

import (
	"github.com/jesserahman/goLangPracticeProject/domain"
	"github.com/jesserahman/goLangPracticeProject/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetCustomerById(string) (*domain.Customer, *errs.AppError)
	GetCustomersByStatus(string) ([]domain.Customer, *errs.AppError)
}

type CustomerServiceImpl struct {
	repository domain.CustomerRepository
}

func (service CustomerServiceImpl) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	return service.repository.FindAll()
}

func (service CustomerServiceImpl) GetCustomerById(id string) (*domain.Customer, *errs.AppError) {
	return service.repository.ById(id)
}

func (service CustomerServiceImpl) GetCustomersByStatus(status string) ([]domain.Customer, *errs.AppError) {
	return service.repository.ByStatus(status)
}

func NewCustomerService(repo domain.CustomerRepository) CustomerServiceImpl {
	return CustomerServiceImpl{repo}
}
