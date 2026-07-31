package dto

type RegisterResponse struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}