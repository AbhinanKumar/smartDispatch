package service

import "github.com/AbhinanKumar/smartDispatch/internal.dto"

type AuthService interface{
	Register(req dto.RegisterRequest) (*dto.RegisterResponse, error)
}