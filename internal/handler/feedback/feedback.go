package feedback

import (
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/golang_kantin/internal/api"
	"github.com/onainadapdap1/golang_kantin/internal/service/feedback"
	"github.com/onainadapdap1/golang_kantin/models"
)

type FeedbackHandler interface {
	CreateFeedback(c *gin.Context)
	GetAllMyFeedback(c *gin.Context)
}

type feedbackHandler struct {
	service feedback.FeedbackService
}

func NewFeedbackHandler(service feedback.FeedbackService) FeedbackHandler {
	return &feedbackHandler{service: service}
}

func (h *feedbackHandler) CreateFeedback(c *gin.Context) {
	var feedback api.CreateFeedbackInput
	if err := c.ShouldBind(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error brewuu": err.Error()})
		return
	}

	log.Println("feedback input : ", feedback)
	log.Println("tipe user id : ", reflect.TypeOf(feedback.UserID))


	parsedTime, err := time.Parse("2006-01-02", feedback.Date)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	feedbackInput := models.Feedback{
		UserID:        feedback.UserID,
		Date:          parsedTime,
		ValueRating:   feedback.ValueRating,
		SubjectReview: feedback.SubjectReview,
		Description:   feedback.SubjectReview,
		File:          feedback.File,
	}

	if err := h.service.CreateFeedback(&feedbackInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create feedback"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Feedback created successfully"})
}

func (h *feedbackHandler) GetAllMyFeedback(c *gin.Context) {
	feedbacks, err := h.service.GetAllMyFeedback()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": feedbacks})
}