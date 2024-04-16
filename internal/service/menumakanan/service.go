package menumakanan

import (
	"github.com/onainadapdap1/golang_kantin/internal/repository/menumakanan"
	"github.com/onainadapdap1/golang_kantin/models"
)

type MenuMakananService interface {
	CreateMenuMakanan(menu models.MenuMakanan) error
	GetAllMenuMakanan() ([]models.MenuMakanan, error)
}

type menuMakananRepository struct {
	repo menumakanan.MenuMakananRepository
}

func NewMenuMakananServ(repo menumakanan.MenuMakananRepository) MenuMakananService {
	return &menuMakananRepository{repo: repo}
}

func (s *menuMakananRepository) CreateMenuMakanan(menu models.MenuMakanan) error {
	return s.repo.CreateMenuMakanan(menu)
}

func (s *menuMakananRepository) GetAllMenuMakanan() ([]models.MenuMakanan, error) {
	menusMakanans, err := s.repo.GetAllMenuMakanan()
	if err != nil {
		return nil, err
	}
	return menusMakanans, nil
}