package internal

import (
	"log"

	"go.uber.org/dig"
)

func RunDig() {
	log.Println("Running dig example")

	c := dig.New()

	deps := []interface{}{
		NewUserService,
		NewAuthenticationService,
		NewHandlerV0,
		NewHandlerV1,
	}

	for _, dep := range deps {
		if err := c.Provide(dep); err != nil {
			panic(err)
		}
	}

	var handler HandlerV0
	if err := c.Invoke(func(h *HandlerV0) {
		handler = *h
	}); err != nil {
		panic(err)
	}
	handler.Handle()

	var handlerV1 HandlerV1
	if err := c.Invoke(func(h *HandlerV1) {
		handlerV1 = *h
	}); err != nil {
		panic(err)
	}

	handlerV1.Handle()
}
