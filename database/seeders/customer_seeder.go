package seeders

import (
	"customer-service/domain/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RunCustomerSeeder(db *gorm.DB) {
	customer := models.Customer{
		ID:     1,
		Name:   "Muhammad Aditya",
		Email:  "m.aditya3232@gmail.com",
		Status: "ACTIVE",
	}

	err := db.FirstOrCreate(&customer, models.Customer{Email: customer.Email}).Error
	if err != nil {
		logrus.Errorf("failed to seed customer: %v", err)
		panic(err)
	}

	logrus.Infof("customer %s successfully seeded", customer.Email)
}
