package allergyreport

import (
	"github.com/onainadapdap1/golang_kantin/internal/repository/allergyreport"
	"github.com/onainadapdap1/golang_kantin/models"
)

type AllergyReportServ interface {
	CheckIsReportExist(userID uint) bool
	CreateAllergyReport(report models.AllergyReport) (models.AllergyReport, error)
	GetAllAllergyReportByUserId(userID uint) ([]models.AllergyReport, error)
}

type allergyreportServ struct {
	repo allergyreport.AllergyReportRepo
}

func NewAllergyReportServ(repo allergyreport.AllergyReportRepo) AllergyReportServ {
	return &allergyreportServ{repo: repo}
}

func (s *allergyreportServ) CheckIsReportExist(userID uint) bool {
	isExist := s.repo.CheckIsReportExist(userID)
	if isExist == 1 {
		return true
	}
	return false
}

func (s *allergyreportServ) CreateAllergyReport(report models.AllergyReport) (models.AllergyReport, error) {
	if err := s.repo.CreateReportAllergy(&report); err != nil {
		return models.AllergyReport{}, err
	}
	return report, nil
	// return s.repo.CreateReportAllergy(*report)
}

func (s *allergyreportServ) GetAllAllergyReportByUserId(userID uint) ([]models.AllergyReport, error) {
	allergyReports, err := s.repo.GetAllAllergyReportByUserId(userID)
	if err != nil {
		return []models.AllergyReport{}, err
	}

	return allergyReports, nil
}
