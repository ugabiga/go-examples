package internal

import (
	"log"

	"go.uber.org/fx"
)

func RunFx() {
	log.Println("Running dig example")

	var handlerV0 HandlerV0
	var handlerV1 HandlerV1
	fx.New(
		fx.Provide(
			NewUserService,
			NewAuthenticationService,
			NewHandlerV0,
			NewHandlerV1,
		),
		fx.Invoke(func(lc fx.Lifecycle, v0 *HandlerV0, v1 *HandlerV1) {
			handlerV0 = *v0
			handlerV1 = *v1
		}),
	)

	handlerV0.Handle()
	handlerV1.Handle()
}
