package models

import (
	"time"
)

type Pengumuman struct {
	ID               uint `gorm:"primarykey"`
	TanggalBerakhir  time.Time
	TanggalPembuatan time.Time
	Deskripsi        string
	Published        bool `gorm:"default:false"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
