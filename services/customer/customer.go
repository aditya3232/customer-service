package services

import (
	"context"
	errConstant "customer-service/constants/error"
	"customer-service/domain/dto"
	"customer-service/repositories"
)

type CustomerService struct {
	repository repositories.IRepositoryRegistry
}

type ICustomerService interface {
	FindByID(context.Context, int) (*dto.CustomerResponse, error)
	Create(context.Context, *dto.CustomerRequest) (*dto.CustomerResponse, error)
	FindAllWithoutPagination(context.Context) ([]dto.CustomerResponse, error)
}

func NewCustomerService(repository repositories.IRepositoryRegistry) ICustomerService {
	return &CustomerService{repository: repository}
}

func (s *CustomerService) FindByID(ctx context.Context, id int) (*dto.CustomerResponse, error) {
	customer, err := s.repository.GetCustomer().FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := dto.CustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		Status:    customer.Status,
		CreatedAt: customer.CreatedAt,
	}

	return &response, nil
}

func (s *CustomerService) IsEmailExist(ctx context.Context, email string) bool {
	customer, err := s.repository.GetCustomer().FindByEmail(ctx, email)
	if err != nil {
		return false
	}

	if customer != nil {
		return true
	}

	return false
}

func (s *CustomerService) Create(ctx context.Context, req *dto.CustomerRequest) (*dto.CustomerResponse, error) {
	if s.IsEmailExist(ctx, req.Email) {
		return nil, errConstant.ErrEmailExist
	}

	if !req.Status.IsValid() {
		return nil, errConstant.ErrStatus
	}

	customer, err := s.repository.GetCustomer().Create(ctx, &dto.CustomerRequest{
		Name:   req.Name,
		Email:  req.Email,
		Status: req.Status,
	})

	if err != nil {
		return nil, err
	}

	response := &dto.CustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		Status:    req.Status,
		CreatedAt: customer.CreatedAt,
	}

	return response, nil
}

func (s *CustomerService) FindAllWithoutPagination(ctx context.Context) ([]dto.CustomerResponse, error) {
	customers, err := s.repository.GetCustomer().FindAllWithoutPagination(ctx)
	if err != nil {
		return nil, err
	}

	customerResults := make([]dto.CustomerResponse, 0, len(customers))
	for _, customer := range customers {
		customerResults = append(customerResults, dto.CustomerResponse{
			ID:        customer.ID,
			Name:      customer.Name,
			Email:     customer.Email,
			Status:    customer.Status,
			CreatedAt: customer.CreatedAt,
		})
	}

	return customerResults, nil
}
