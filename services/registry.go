package services

import (
	"customer-service/repositories"
	customerService "customer-service/services/customer"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
}

type IServiceRegistry interface {
	GetCustomer() customerService.ICustomerService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry) IServiceRegistry {
	return &Registry{repository: repository}
}

func (r *Registry) GetCustomer() customerService.ICustomerService {
	return customerService.NewCustomerService(r.repository)
}
