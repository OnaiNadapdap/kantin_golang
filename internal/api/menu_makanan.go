package api

type CreateMenuMakananInput struct {
	TanggalMakan string `form:"tanggal_makan" json:"tanggal_makan" binding:"required"`
	MenuPagi     string `form:"menu_pagi" json:"menu_pagi" binding:"required"`
	MenuSiang    string `form:"menu_siang" json:"menu_siang" binding:"required"`
	MenuMalam    string `form:"menu_malam" json:"menu_malam" binding:"required"`
	Foto1        string `form:"foto_1" json:"foto_1"`
	Foto2        string `form:"foto_2" json:"foto_2"`
	Foto3        string `form:"foto_3" json:"foto_3"`
}

// type MenuMakanan struct {
//     ID           uint      `gorm:"primaryKey"`
//     TanggalMakan time.Time `gorm:"column:tanggal_makan"`
//     MenuPagi     string    `gorm:"column:menu_pagi"`
//     MenuSiang    string    `gorm:"column:menu_siang"`
//     MenuMalam    string    `gorm:"column:menu_malam"`
//     Foto1        string    `gorm:"column:foto1"`
//     Foto2        string    `gorm:"column:foto2"`
//     Foto3        string    `gorm:"column:foto3"`
//     CreatedAt    time.Time `gorm:"column:created_at"`
//     UpdatedAt    time.Time `gorm:"column:updated_at"`
// }
