package allergyreport

import (
	"log"

	"github.com/onainadapdap1/golang_kantin/models"
	"gorm.io/gorm"
)

type AllergyReportRepo interface {
	CheckIsReportExist(userID uint) int
	CreateReportAllergy(report *models.AllergyReport) error
	GetAllAllergyReportByUserId(userID uint) ([]models.AllergyReport, error)
}

type allergyreportRepo struct {
	db *gorm.DB
}

func NewAllergyReportRepo(db *gorm.DB) AllergyReportRepo {
	return &allergyreportRepo{db: db}
}

func (r *allergyreportRepo) CheckIsReportExist(userID uint) int {
	var report models.AllergyReport
	log.Println("hello user id alergy : ", userID)
	if err := r.db.Where("user_id", userID).Where("approved", 0).First(&report).Error; err != nil {
		return 0
	}

	return 1
}

func (r *allergyreportRepo) CreateReportAllergy(report *models.AllergyReport) error {
	return r.db.Create(&report).Error
}

func (r *allergyreportRepo) GetAllAllergyReportByUserId(userID uint) ([]models.AllergyReport, error) {
	var allergyReports []models.AllergyReport
	if err := r.db.Where("user_id", userID).Preload("User").Find(&allergyReports).Error; err != nil {
		return nil, err
	}

	return allergyReports, nil
}