package allergyreport

import (
	"github.com/onainadapdap1/golang_kantin/models"
	"gorm.io/gorm"
)

type AllergyReportRepo interface {
	CheckIsReportExist(userID uint) (bool)
	CreateReportAllergy(report models.AllergyReport) error
}

type allergyreportRepo struct {
	db *gorm.DB
}

func NewAllergyReportRepo(db *gorm.DB) AllergyReportRepo {
	return &allergyreportRepo{db: db}
}

func (r *allergyreportRepo) CheckIsReportExist(userID uint) (bool) {
	var report models.AllergyReport
	if err := r.db.Where("user_id", userID).Where("approved", 0).First(&report).Error; err != nil {
		return false
	}

	return true
}

func (r *allergyreportRepo) CreateReportAllergy(report models.AllergyReport) error {
	return r.db.Create(&report).Error
}