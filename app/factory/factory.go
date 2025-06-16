package factory

import "souzalambdago/service"

type RouteKey struct {
	Method string
	Path   string
}

type Factory struct {
	Routes map[RouteKey]service.Service
}

func NewFactory() *Factory {
	return &Factory{
		Routes: map[RouteKey]service.Service{
			{"GET", "/payment"}:  service.NewGetPaymentService(),
			{"POST", "/payment"}: service.NewCreatePaymentService(),
		},
	}
}

func (f *Factory) GetService(method, path string) (service.Service, bool) {
	svc, ok := f.Routes[RouteKey{method, path}]
	return svc, ok
}
