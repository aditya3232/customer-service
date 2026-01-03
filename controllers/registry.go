package controllers

import (
	customerController "customer-service/controllers/customer"
	"customer-service/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetCustomer() customerController.ICustomerController
}

func NewControllerregistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}

func (r *Registry) GetCustomer() customerController.ICustomerController {
	return customerController.NewCustomerController(r.service)
}
