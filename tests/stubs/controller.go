package stubs

import "github.com/dougmendes/gondalf/pkg/infra/controller"

type ControllerStub struct {
	AddRoutesFunc func(controller.Router)
}

func (c ControllerStub) AddRoutes(routes controller.Router) {
	c.AddRoutesFunc(routes)
}
