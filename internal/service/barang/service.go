package barang

import (
	"fmt"

	"github.com/onainadapdap1/golang_kantin/internal/repository/barang"
	"github.com/onainadapdap1/golang_kantin/models"
)

type BarangService interface {
	CreateBarang(barang *models.Barang) error
	ShowBarang(id uint) error
	HideBarang(id uint) error
	GetPengumuman(page int, perPage int) ([]models.Barang, error)
}

type barangService struct {
	repo barang.BarangRepository
}

func NewBarangService(repo  barang.BarangRepository) BarangService {
	return &barangService{repo: repo}
}

func (s *barangService) CreateBarang(barang *models.Barang) error {
	return s.repo.CreateBarang(barang)
}

func (s *barangService) ShowBarang(id uint) error {
	// barang := models.Barang{}
	barang, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	err = s.repo.ShowBarang(barang)
	if err != nil {
		return err
	}

	return nil
}

func (s *barangService) HideBarang(id uint) error {
	barang, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	err = s.repo.HideBarang(barang)
	if err != nil {
		return err
	}

	return nil
}

func (s *barangService) GetPengumuman(page int, perPage int) ([]models.Barang, error) {
	fmt.Println("logging here serv")
    pengumuman, err := s.repo.GetPengumuman(page, perPage)
    if err != nil {
        return nil, err
    }

    return pengumuman, nil
}