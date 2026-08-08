package service

import "github.com/AbhinanKumar/smart-dispatch/internal/dto"

type CustomerService interface {
	Create(req dto.CreateCustomerRequest) (*dto.CustomerResponse, error)
	GetByID(id uint) (*dto.CustomerResponse, error)
	GetAll() ([]dto.CustomerResponse, error)
	Update(id uint, req dto.CreateCustomerRequest) (*dto.CustomerResponse, error)
}
