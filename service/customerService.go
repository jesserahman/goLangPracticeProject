package service

import "github.com/jesserahman/goLangPracticeProject/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

type CustomerServiceImpl struct {
	repository domain.CustomerRepository
}

func (service CustomerServiceImpl) GetAllCustomers() ([]domain.Customer, error) {
	return service.repository.FindAll()
}

func (service CustomerServiceImpl) GetCustomer(id string) (*domain.Customer, error) {
	return service.repository.ById(id)
}

func NewCustomerService(repo domain.CustomerRepository) CustomerServiceImpl {
	return CustomerServiceImpl{repo}
}
