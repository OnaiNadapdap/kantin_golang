package api

type CreatePengumumanInput struct {
	TanggalBerakhir     string `gorm:"omitempty" form:"tanggal_berakhir" json:"tanggal_berakhir"`
	Deskripsi     string `gorm:"omitempty" form:"deskripsi" json:"deskripsi"`
	// User           model.User
}