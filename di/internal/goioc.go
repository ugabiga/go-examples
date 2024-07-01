package internal

import (
	"log"
	"reflect"

	"github.com/goioc/di"
)

func RunGoIoC() {
	log.Println("Running goioc example")

	_, _ = di.RegisterBean("srv", reflect.TypeOf((*UserService)(nil)))
	_, _ = di.RegisterBean("handler", reflect.TypeOf((*HandlerV0)(nil)))
	_ = di.InitializeContainer()

	handler := di.GetInstance("handler").(*HandlerV0)
	handler.Handle()
}
