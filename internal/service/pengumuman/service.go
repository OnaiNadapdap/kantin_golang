package pengumuman

import (
	"time"

	"github.com/onainadapdap1/golang_kantin/internal/repository/pengumuman"
	"github.com/onainadapdap1/golang_kantin/models"
)


type PengumumanService interface {
	CreatePengumuman(pengumuman models.Pengumuman) (models.Pengumuman, error)
}

type pengumumanService struct {
	repo pengumuman.PengumumanRepository
}

func NewPengumumanService(repo pengumuman.PengumumanRepository) PengumumanService {
	return &pengumumanService{repo}
}

func (s *pengumumanService) CreatePengumuman(pengumuman models.Pengumuman) (models.Pengumuman, error) {
	pengumuman.TanggalPembuatan = time.Now()
	if err := s.repo.CreatePengumuman(&pengumuman); err != nil {
		return models.Pengumuman{}, err
	}
	return pengumuman, nil
}