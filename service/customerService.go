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
	CreateNewCustomer(dto.NewCustomerRequest) (*dto.NewCustomerResponse, *errs.AppError)
	UpdateCustomer(customer domain.Customer) (*dto.CustomerResponse, *errs.AppError)
	DeleteCustomer(customerId string) *errs.AppError
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
		customerDto := customer.ToCustomerResponseDto()
		customersDto = append(customersDto, *customerDto)
	}
	return customersDto, nil
}

func (service CustomerServiceImpl) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := service.repository.ById(id)
	if err != nil {
		return nil, err
	}
	return customer.ToCustomerResponseDto(), nil
}

func (service CustomerServiceImpl) GetCustomersByStatus(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := service.repository.ByStatus(status)
	if err != nil {
		return nil, err
	}
	var customersDto []dto.CustomerResponse
	for _, customer := range customers {
		customerDto := customer.ToCustomerResponseDto()
		customersDto = append(customersDto, *customerDto)
	}
	return customersDto, nil
}

func (service CustomerServiceImpl) CreateNewCustomer(newCustomerRequestDto dto.NewCustomerRequest) (*dto.NewCustomerResponse, *errs.AppError) {
	err := newCustomerRequestDto.Validate()
	if err != nil {
		return nil, err
	}

	customer := domain.Customer{
		Name:        newCustomerRequestDto.Name,
		City:        newCustomerRequestDto.City,
		Zip:         newCustomerRequestDto.ZipCode,
		Status:      newCustomerRequestDto.Status,
		DateOfBirth: newCustomerRequestDto.DateOfBirth,
	}
	response, saveCustomerErr := service.repository.Save(customer)
	if saveCustomerErr != nil {
		return nil, saveCustomerErr
	}

	return response.ToNewCustomerResponseDto(), nil
}

func (service CustomerServiceImpl) UpdateCustomer(customer domain.Customer) (*dto.CustomerResponse, *errs.AppError) {
	response, updateCustomerErr := service.repository.Update(customer)
	if updateCustomerErr != nil {
		return nil, updateCustomerErr
	}

	return response.ToCustomerResponseDto(), nil
}

func (service CustomerServiceImpl) DeleteCustomer(customerId string) *errs.AppError {
	err := service.repository.Delete(customerId)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerService(repo domain.CustomerRepository) CustomerServiceImpl {
	return CustomerServiceImpl{repo}
}
