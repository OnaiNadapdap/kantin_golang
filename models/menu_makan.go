package models

import "time"

type MenuMakanan struct {
    ID           uint      `gorm:"primaryKey"`
    TanggalMakan time.Time `gorm:"column:tanggal_makan"`
    MenuPagi     string    `gorm:"column:menu_pagi"`
    MenuSiang    string    `gorm:"column:menu_siang"`
    MenuMalam    string    `gorm:"column:menu_malam"`
    Foto1        string    `gorm:"column:foto1"`
    Foto2        string    `gorm:"column:foto2"`
    Foto3        string    `gorm:"column:foto3"`
    CreatedAt    time.Time `gorm:"column:created_at"`
    UpdatedAt    time.Time `gorm:"column:updated_at"`
}

// TableName sets the table name for the model
func (MenuMakanan) TableName() string {
    return "menu_makanans"
}