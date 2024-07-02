package internal

func RunWire() {
	handlerV0 := InitializeHandlerV0()
	handlerV0.Handle()

	handlerV1 := InitializeHandlerV1()
	handlerV1.Handle()
}
