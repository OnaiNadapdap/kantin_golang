package allergyreport

import (
	"github.com/onainadapdap1/golang_kantin/internal/repository/allergyreport"
	"github.com/onainadapdap1/golang_kantin/models"
)

type AllergyReportServ interface {
	CheckIsReportExist(userID uint)  (bool)
	CreateAllergyReport(report *models.AllergyReport) error
}

type allergyreportServ struct {
	repo allergyreport.AllergyReportRepo
}

func NewAllergyReportServ(repo allergyreport.AllergyReportRepo) AllergyReportServ {
	return &allergyreportServ{repo: repo}
}

func (s *allergyreportServ) CheckIsReportExist(userID uint)  (bool) {
	return s.repo.CheckIsReportExist(userID)
}

func (s *allergyreportServ) CreateAllergyReport(report *models.AllergyReport) error {
	return s.repo.CreateReportAllergy(*report)
}