package pengumuman

import (
	"github.com/onainadapdap1/golang_kantin/models"
	"gorm.io/gorm"
)

type PengumumanRepository interface {
	CreatePengumuman(pengumuman *models.Pengumuman) error
}

type pengumumanRepository struct {
	db *gorm.DB
}

func NewPengumumanRepository(db *gorm.DB) PengumumanRepository {
	return &pengumumanRepository{db}
}

func (r *pengumumanRepository) CreatePengumuman(pengumuman *models.Pengumuman) error {
	return r.db.Create(pengumuman).Error
}
