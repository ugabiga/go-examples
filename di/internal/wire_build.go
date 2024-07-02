//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
)

func InitializeHandlerV0() *HandlerV0 {
	wire.Build(NewHandlerV0, NewUserService)
	return &HandlerV0{}
}

func InitializeHandlerV1() *HandlerV1 {
	wire.Build(NewHandlerV1, NewHandlerV1Deps, NewUserService, NewAuthenticationService)
	return &HandlerV1{}
}
