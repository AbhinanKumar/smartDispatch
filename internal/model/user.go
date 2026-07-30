package model

import "time"

type User struct{
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`
	Email string `gorm:"size:255;uniqueIndex;not null`
	PasswordHash string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// gorm.Model            gorm.Model automatically adds: ID, CreatedAt, UpdatedAt, DeletedAt
}