package feedback

import (
	"github.com/onainadapdap1/golang_kantin/internal/repository/feedback"
	"github.com/onainadapdap1/golang_kantin/models"
)

type FeedbackService interface {
	CreateFeedback(feedback *models.Feedback) error
	GetAllFeedback() ([]models.Feedback, error)
	GetAllMyFeedback(userID uint) ([]models.Feedback, error)
}

type feedbackService struct {
	repository feedback.FeedbackRepository
}

func NewFeedbackService(repository feedback.FeedbackRepository) FeedbackService {
	return &feedbackService{repository: repository}
}

func (s *feedbackService) CreateFeedback(feedback *models.Feedback) error {
	return s.repository.CreateFeedback(feedback)
}

func (s *feedbackService) GetAllFeedback() ([]models.Feedback, error) {
	feedbacks, err := s.repository.GetAllFeedback()
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (s *feedbackService) GetAllMyFeedback(userID uint) ([]models.Feedback, error) {
	myFeedbacks, err := s.repository.GetAllMyFeedback(userID)
	if err != nil {
		return nil, err
	}

	return myFeedbacks, nil
}