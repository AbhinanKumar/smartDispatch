package repository

import "github.com/AbhinanKumar/smart-dispatch/internal/model"

type CustomerRepository interface {
	Create(customer *model.Customer) error
	GetByID(id uint) (*model.Customer, error)
	GetAll() ([]model.Customer, error)
	Update(customer *model.Customer) error
	FindByemail(email string) (*model.Customer, error)
}
