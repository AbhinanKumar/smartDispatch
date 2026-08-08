package model

import "time"

type Request struct {
	ID uint

	CustomerID uint

	CustomerNameSnapshot  string
	CustomerPhoneSnapshot string

	Title       string
	Description string
	Status      string

	CreatedAt time.Time
	UpdatedAt time.Time
}
