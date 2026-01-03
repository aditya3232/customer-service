package models

import (
	"customer-service/constants"
	"time"
)

type Customer struct {
	ID        int                            `gorm:"primaryKey;autoIncrement"`
	Name      string                         `gorm:"type:varchar(100);not null"`
	Email     string                         `gorm:"type:varchar(100);not null"`
	Status    constants.CustomerStatusString `gorm:"not null"`
	CreatedAt *time.Time
}
