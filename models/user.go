package models

import (
	"time"
)

type User struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	RoleID          uint      `json:"role_id"`
	Role            Role      `gorm:"foreignKey:RoleID"`
	Nim             string    `json:"nim"`
	Name            string    `json:"name"`
	Prodi           string    `json:"prodi"`
	Angkatan        string    `json:"angkatan"`
	Asrama          string    `json:"asrama"`
	Email           string    `gorm:"unique" json:"email"`
	EmailVerifiedAt time.Time `gorm:"default:null" json:"email_verified_ata"`
	Password        string    `json:"password"`
	RememberToken   string    `json:"remember_token"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	AllergyReport   []AllergyReport
}
