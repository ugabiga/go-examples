package internal

import (
	"log"

	"go.uber.org/dig"
)

type HandlerV0 struct {
	srv *UserService `di.inject:"srv"`
}

func NewHandlerV0(srv *UserService) *HandlerV0 {
	return &HandlerV0{srv: srv}
}

func (h HandlerV0) Handle() {
	log.Println("Handler v0 is handling")
	h.srv.Run()
}

type HandlerV1Deps struct {
	dig.In

	UserSrv *UserService
	AuthSrv *AuthenticationService
}

// NewHandlerV1Deps for wire
func NewHandlerV1Deps(userSrv *UserService, authSrv *AuthenticationService) HandlerV1Deps {
	return HandlerV1Deps{
		UserSrv: userSrv,
		AuthSrv: authSrv,
	}
}

type HandlerV1 struct {
	userSrv *UserService
	authSrv *AuthenticationService
}

func NewHandlerV1(deps HandlerV1Deps) *HandlerV1 {
	return &HandlerV1{
		userSrv: deps.UserSrv,
		authSrv: deps.AuthSrv,
	}
}

func (h HandlerV1) Handle() {
	log.Println("Handler v1 is handling")
	h.userSrv.Run()
	h.authSrv.Run()
}
