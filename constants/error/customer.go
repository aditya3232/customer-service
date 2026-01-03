package error

import "errors"

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrEmailExist       = errors.New("email already exist")
	ErrStatus           = errors.New("status must be ACTIVE or INACTIVE")
)

var CustomerErrors = []error{
	ErrCustomerNotFound,
	ErrEmailExist,
	ErrStatus,
}
