package entity

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Name cannot blank"`
	Email      string
	EmployeeID string `valid:"matches(^[JMS]+\\d{8}$)~EmployeeID invalid format"`
}
