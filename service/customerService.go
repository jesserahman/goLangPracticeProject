package service

import (
	"github.com/jesserahman/goLangPracticeProject/domain"
	"github.com/jesserahman/goLangPracticeProject/dto"
	"github.com/jesserahman/goLangPracticeProject/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppError)
	GetCustomersByStatus(string) ([]dto.CustomerResponse, *errs.AppError)
}

type CustomerServiceImpl struct {
	repository domain.CustomerRepository
}

func (service CustomerServiceImpl) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}
	var customersDto []dto.CustomerResponse
	for _, customer := range customers {
		customerDto := customer.ToDto()
		customersDto = append(customersDto, *customerDto)
	}
	return customersDto, nil
}

func (service CustomerServiceImpl) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := service.repository.ById(id)
	if err != nil {
		return nil, err
	}
	return customer.ToDto(), nil
}

func (service CustomerServiceImpl) GetCustomersByStatus(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := service.repository.ByStatus(status)
	if err != nil {
		return nil, err
	}
	var customersDto []dto.CustomerResponse
	for _, customer := range customers {
		customerDto := customer.ToDto()
		customersDto = append(customersDto, *customerDto)
	}
	return customersDto, nil
}

func NewCustomerService(repo domain.CustomerRepository) CustomerServiceImpl {
	return CustomerServiceImpl{repo}
}
