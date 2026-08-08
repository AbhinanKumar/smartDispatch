package model

import "time"

type Customer struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Phone     string `gorm:"size:20;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
