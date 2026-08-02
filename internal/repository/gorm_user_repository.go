package repository

import (
	"gorm.io/gorm"
	"github.com/AbhinanKumar/smart-dispatch/internal/model"
)

type GormUserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &GormUserRepository{
		db: db,
	}
}

func (r *GormUserRepository) FindByemail(email string) (*model.User, error){
	var user model.User		// use the zero-value struct and return its address later. This avoids allocating the pointer immediately

	err := r.db.Where("email = ?",email).First(&user).Error			//First(&user) -> GORM must populate the struct.
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil, nil
		}
		return nil, err
	}
	return &user, nil				//service needs the populated struct. Returning a pointer avoids copying the entire struct.
}

func (r *GormUserRepository) Create(user *model.User) error{
	return r.db.Create(user).Error
}