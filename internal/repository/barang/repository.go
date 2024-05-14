package barang

import (
	"fmt"
	"log"
	"time"

	"github.com/onainadapdap1/golang_kantin/models"
	"gorm.io/gorm"
)

type BarangRepository interface {
	CreateBarang(barang *models.Barang) error
	ShowBarang(barang models.Barang) error
	FindByID(id uint) (models.Barang, error)
	HideBarang(barang models.Barang) error
	GetPengumuman(page int, perPage int) ([]models.Barang, error)
}

type barangRepository struct {
	db *gorm.DB
}

func NewBarangRepository(db *gorm.DB) BarangRepository {
	return &barangRepository{db: db}
}

func (r *barangRepository) CreateBarang(barang *models.Barang) error {
	return r.db.Create(barang).Error
}

func (r *barangRepository) FindByID(id uint) (models.Barang, error) {
	log.Println("id in findbyid : ", id)
	tx := r.db.Begin()
	var barang models.Barang
	if err := tx.Debug().First(&barang, id).Error; err != nil {
		return models.Barang{}, err
	}
	log.Println("barang in repo is : ", barang)
	return barang, nil
}

func (r *barangRepository) ShowBarang(barang models.Barang) error {
	tx := r.db.Begin()
	barang.Showed = 1
	if err := tx.Debug().Save(&barang).Error; err != nil {
		return err
	}

	tx.Commit()
	return nil

}

func (r *barangRepository) HideBarang(barang models.Barang) error {
	tx := r.db.Begin()
	barang.Showed = 0
	if err := tx.Debug().Save(&barang).Error; err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (r *barangRepository) GetPengumuman(page int, perPage int) ([]models.Barang, error) {
	fmt.Println("loggin here repo")
	var pengumuman []models.Barang
	today := time.Now().Format("2006-01-02")
	tx := r.db.Begin()

	if err := tx.Debug().Where("showed = ?", 1).
		Where("expiry_date >= ?", today).
		Order("kategori").
		Order("expiry_date desc").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&pengumuman).Error; err != nil {
		return nil, err
	}

	return pengumuman, nil
}
