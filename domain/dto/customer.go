package dto

import (
	"customer-service/constants"
	"time"
)

type CustomerRequest struct {
	Name   string                         `json:"name" validate:"required"`
	Email  string                         `json:"email" validate:"required,email"`
	Status constants.CustomerStatusString `json:"status"`
}

type CustomerResponse struct {
	ID        int                            `json:"id"`
	Name      string                         `json:"name"`
	Email     string                         `json:"email"`
	Status    constants.CustomerStatusString `json:"status"`
	CreatedAt *time.Time                     `json:"created_at"`
}
