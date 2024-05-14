package api

type AllergyReportInput struct {
	// UserID    uint   `json:"user_id" form:"user_id"`
	Allergies string `json:"allergies" form:"allergies"`
	File      string `json:"file" form:"file"`
}

// type AllergyReport struct {
// 	ID              uint `gorm:"primaryKey"`
// 	UserID          uint
// 	Allergies       string
// 	File            string
// 	Approved        bool
// 	AlasanPenolakan string
// 	CreatedAt       time.Time
// 	UpdatedAt       time.Time
// }
