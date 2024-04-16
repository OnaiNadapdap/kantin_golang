package models

import (
	"time"
)

type Feedback struct {
	ID            uint `gorm:"primarykey"`
	UserID        uint
	Date          time.Time
	ValueRating   string
	SubjectReview string
	Description   string
	File          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
