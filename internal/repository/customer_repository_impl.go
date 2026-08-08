package repository

import (
	"github.com/AbhinanKumar/smart-dispatch/internal/model"
	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepo{
		db: db,
	}
}

func (r *customerRepo) Create(customer *model.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepo) GetByID(id uint) (*model.Customer, error) {
	var customer model.Customer

	err := r.db.First(&customer, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepo) GetAll() ([]model.Customer, error) {
	var customers []model.Customer

	err := r.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *customerRepo) FindByemail(email string) (*model.Customer, error) {
	var customer model.Customer

	err := r.db.Where("email = ?", email).First(&customer).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Why return (nil, nil)? --> Database worked. Customer simply doesn't exist. That is not an error.
	}

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepo) Update(customer *model.Customer) error {
	return r.db.Save(customer).Error
}
