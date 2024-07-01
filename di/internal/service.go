package internal

import "log"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (srv UserService) Run() {
	log.Println("User Service is running")
}

type AuthenticationService struct {
}

func NewAuthenticationService() *AuthenticationService {
	return &AuthenticationService{}
}

func (srv AuthenticationService) Run() {
	log.Println("Authentication Service is running")
}
