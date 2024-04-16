package pengumuman

import (
	"github.com/onainadapdap1/golang_kantin/models"
	"gorm.io/gorm"
)

type PengumumanRepository interface {
	CreatePengumuman(pengumuman *models.Pengumuman) error
	GetAllPengumuman() ([]models.Pengumuman, error)
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

func (r *pengumumanRepository) GetAllPengumuman() ([]models.Pengumuman, error) {
	tx := r.db.Begin()
	var pengumuman []models.Pengumuman
	if err := tx.Debug().Find(&pengumuman).Error; err != nil {
		return nil, err
	}

	return pengumuman, nil
}
