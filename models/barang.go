package models

import "time"

type Barang struct {
	ID          uint      `gorm:"primaryKey"`
	Kategori    string    `gorm:"not null"`
	UserID      uint      `gorm:"not null"`
	// User        User      // relasi ke tabel users
	User        User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"user"`
	Name        string    // nullable di Laravel sama dengan string kosong di Golang
	Description string    // nullable di Laravel sama dengan string kosong di Golang
	ExpiryDate  time.Time // nullable di Laravel sama dengan pointer di Golang
	File        string    // nullable di Laravel sama dengan string kosong di Golang
	Showed      int       `gorm:"default:0"` // diubah menjadi int, sesuai dengan tinyInteger di Laravel
	CreatedAt   time.Time // tipe default di Golang, sesuai dengan timestamps di Laravel
	UpdatedAt   time.Time // tipe default di Golang, sesuai dengan timestamps di Laravel
}
