package repositories

import (
	"context"
	errWrap "customer-service/common/error"
	errConstant "customer-service/constants/error"
	"customer-service/domain/dto"
	"customer-service/domain/models"
	"errors"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

type ICustomerRepository interface {
	FindByID(context.Context, int) (*models.Customer, error)
	FindByEmail(context.Context, string) (*models.Customer, error)
	Create(context.Context, *dto.CustomerRequest) (*models.Customer, error)
	FindAllWithoutPagination(context.Context) ([]models.Customer, error)
}

func NewCustomerRepository(db *gorm.DB) ICustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) FindByID(ctx context.Context, id int) (*models.Customer, error) {
	var customer models.Customer

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConstant.ErrCustomerNotFound
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &customer, nil
}

func (r *CustomerRepository) Create(ctx context.Context, req *dto.CustomerRequest) (*models.Customer, error) {
	customer := models.Customer{
		Name:   req.Name,
		Email:  req.Email,
		Status: req.Status,
	}

	err := r.db.WithContext(ctx).Create(&customer).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &customer, nil
}

func (r *CustomerRepository) FindAllWithoutPagination(ctx context.Context) ([]models.Customer, error) {
	var customers []models.Customer

	err := r.db.WithContext(ctx).Find(&customers).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return customers, nil
}

func (r *CustomerRepository) FindByEmail(ctx context.Context, email string) (*models.Customer, error) {
	var customer models.Customer

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConstant.ErrCustomerNotFound
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &customer, nil
}
