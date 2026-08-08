package service

import (
	"errors"
	"github.com/AbhinanKumar/smart-dispatch/internal/dto"
	"github.com/AbhinanKumar/smart-dispatch/internal/model"
	"github.com/AbhinanKumar/smart-dispatch/internal/repository"
	"strings"
)

type custService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &custService{
		custRepo: repo,
	}
}

func (s *custService) Create(req dto.CreateCustomerRequest) (*dto.CustomerResponse, error) {
	if err := validateCustomerRequest(req); err != nil {
		return nil, err
	}

	email := strings.ToLower(strings.TrimSpace(req.Email))
	name := strings.TrimSpace(req.Name)
	phone := strings.TrimSpace(req.Phone)

	existedCustomer, err := s.custRepo.FindByemail(email)
	if err != nil {
		return nil, err
	}

	if existedCustomer != nil {
		return nil, errors.New("customer with this email already exists")
	}

	customer := &model.Customer{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	if err := s.custRepo.Create(customer); err != nil {
		return nil, err
	}
	return &dto.CustomerResponse{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
		Phone: customer.Phone,
	}, nil
}

func validateCustomerRequest(req dto.CreateCustomerRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(req.Phone) == "" {
		return errors.New("phone number is required")
	}
	return nil
}

func (s *custService) GetByID(id uint) (*dto.CustomerResponse, error) {
	cust, err := s.custRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if cust == nil {
		return nil, errors.New("customer not found")
	}

	return &dto.CustomerResponse{
		ID:    cust.ID,
		Name:  cust.Name,
		Email: cust.Email,
		Phone: cust.Phone,
	}, nil
}

func (s *custService) GetAll() ([]dto.CustomerResponse, error) {
	custs, err := s.custRepo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.CustomerResponse, 0, len(custs))

	for _, customer := range custs {
		responses = append(responses, dto.CustomerResponse{
			ID:    customer.ID,
			Name:  customer.Name,
			Email: customer.Email,
			Phone: customer.Phone,
		})
	}

	return responses, nil
}

func (s *custService) Update(id uint, req dto.CreateCustomerRequest) (*dto.CustomerResponse, error) {
	if err := validateCustomerRequest(req); err != nil {
		return nil, err
	}

	cust, err := s.custRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if cust == nil {
		return nil, errors.New("customer not found")
	}

	cust.Name = strings.TrimSpace(req.Name)
	cust.Email = strings.ToLower(strings.TrimSpace(req.Email))
	cust.Phone = strings.TrimSpace(req.Phone)

	if err := s.custRepo.Update(cust); err != nil {
		return nil, err
	}

	return &dto.CustomerResponse{
		ID:    cust.ID,
		Name:  cust.Name,
		Email: cust.Email,
		Phone: cust.Phone,
	}, nil
}
