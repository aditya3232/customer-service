package repositories

import (
	customerRepo "customer-service/repositories/customer"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetCustomer() customerRepo.ICustomerRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Registry{db: db}
}

func (r *Registry) GetCustomer() customerRepo.ICustomerRepository {
	return customerRepo.NewCustomerRepository(r.db)
}
