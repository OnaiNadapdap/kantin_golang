package models

import (
	"time"
)

type AllergyReport struct {
	ID              uint `gorm:"primaryKey"`
	UserID          uint
	Allergies       string
	File            string
	Approved        bool
	AlasanPenolakan string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	User            User
}
