package service

import "github.com/AbhinanKumar/smart-dispatch/internal/dto"

type AuthService interface{
	Register(req dto.RegisterRequest) (*dto.RegisterResponse, error)
}