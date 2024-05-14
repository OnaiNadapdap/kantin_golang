package feedback

import (
	"github.com/onainadapdap1/golang_kantin/models"
	"gorm.io/gorm"
)

type FeedbackRepository interface {
	CreateFeedback(feedback *models.Feedback) error
	GetAllFeedback() ([]models.Feedback, error)
	GetAllMyFeedback(userID uint) ([]models.Feedback, error)
	
}

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepository{db: db}
}

func (r *feedbackRepository) CreateFeedback(feedback *models.Feedback) error {
	return r.db.Create(feedback).Error
}

func (r *feedbackRepository) GetAllFeedback() ([]models.Feedback, error) {
	tx := r.db.Begin()
	var feedbacks []models.Feedback
	if err := tx.Debug().Find(&feedbacks).Error; err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func (r *feedbackRepository) GetAllMyFeedback(userID uint) ([]models.Feedback, error) {
	tx := r.db.Begin()
	var feedbacks []models.Feedback
	if err := tx.Debug().Where("user_id", userID).Find(&feedbacks).Error; err != nil {
		return nil, err
	}
	return feedbacks, nil
}