package api

type CreateInputBarang struct {
	// UserID      uint   `json:"user_id" form:"user_id"`
	Kategori    string `json:"kategori" form:"kategori" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	ExpiryDate  string `json:"expiry_date" form:"expiry_date"`
	File        string `json:"file" form:"file" binding:"required"`
}

// type CreateFeedbackInput struct {
// 	// User           model.User
// 	UserID        uint   `gorm:"omitempty" form:"user_id" json:"user_id"`
// 	Date          string `gorm:"omitempty" form:"date" json:"date"`
// 	ValueRating   string `gorm:"omitempty" form:"value_rating" json:"value_rating"`
// 	SubjectReview string `gorm:"omitempty" form:"subject_review" json:"subject_review"`
// 	Description   string `gorm:"omitempty" form:"description" json:"description"`
// }

// type Barang struct {
//     ID          uint           `gorm:"primaryKey"`
//     Kategori    string         `gorm:"not null"`
//     UserID      uint           `gorm:"not null"`
//     // User        User           // relasi ke tabel users
//     Name        string         // nullable di Laravel sama dengan string kosong di Golang
//     Description string         // nullable di Laravel sama dengan string kosong di Golang
//     ExpiryDate  time.Time     // nullable di Laravel sama dengan pointer di Golang
//     File        string         // nullable di Laravel sama dengan string kosong di Golang
//     Showed      int            `gorm:"default:0"` // diubah menjadi int, sesuai dengan tinyInteger di Laravel
//     CreatedAt   time.Time      // tipe default di Golang, sesuai dengan timestamps di Laravel
//     UpdatedAt   time.Time      // tipe default di Golang, sesuai dengan timestamps di Laravel
// }
