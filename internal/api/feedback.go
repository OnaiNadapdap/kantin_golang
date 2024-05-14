package api

type CreateFeedbackInput struct {
	// User           model.User
	// UserID        uint   `form:"user_id" json:"user_id" binding:"required"`
	Date          string `form:"date" json:"date" binding:"required"`
	ValueRating   string `form:"value_rating" json:"value_rating" binding:"required"`
	SubjectReview string `form:"subject_review" json:"subject_review" binding:"required"`
	Description   string `form:"description" json:"description" binding:"required"`
	File          string `form:"file" json:"file"`
}
