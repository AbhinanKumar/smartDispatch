package repository

import "github.com/AbhinanKumar/smart-dispatch/internal/model"

type UserRepository interface {
	FindByemail(email string) (*model.User, error)
	Create(user *model.User) error
}
