package service

import "github.com/jesserahman/goLangPracticeProject/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type CustomerServiceImpl struct {
	repository domain.CustomerRepository
}

func (service CustomerServiceImpl) GetAllCustomers() ([]domain.Customer, error) {
	return service.repository.FindAll()
}

func NewCustomerService(repo domain.CustomerRepository) CustomerServiceImpl {
	return CustomerServiceImpl{repo}
}
