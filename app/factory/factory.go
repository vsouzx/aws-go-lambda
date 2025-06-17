package factory

import (
	"souzalambdago/repository"
	"souzalambdago/service"
)

type RouteKey struct {
	Method string
	Path   string
}

type Factory struct {
	repository repository.TransactionRepositoryInterface
	routes map[RouteKey]service.Service
}

func NewFactory(repository repository.TransactionRepositoryInterface) *Factory {
	return &Factory{
		repository: repository,
		routes: map[RouteKey]service.Service{
			{"GET", "/payment"}:  service.NewGetPaymentService(repository),
			{"POST", "/payment"}: service.NewCreatePaymentService(repository),
			{"PATCH", "/payment"}: service.NewUpdateTransactionStatusService(repository),
		},
	}
}

func (f *Factory) GetService(method, path string) (service.Service) {
	svc, ok := f.routes[RouteKey{method, path}]
	if !ok {
		panic("Route not found: " + method + " " + path)
	}
	return svc
}
