package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, error)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, error) {
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, appError := s.repo.ById(id)
	if appError != nil {
		return nil, appError
	}

	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repository}
}
