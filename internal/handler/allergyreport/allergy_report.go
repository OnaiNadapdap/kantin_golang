package allergyreport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/golang_kantin/internal/api"
	"github.com/onainadapdap1/golang_kantin/internal/service/allergyreport"
	"github.com/onainadapdap1/golang_kantin/models"
)

type AllergyReportHandler interface {
	CreateAllergyReport(c *gin.Context)
}

type allergyReportHandler struct {
	serv allergyreport.AllergyReportServ
}

func NewAllergyReportHandler(serv allergyreport.AllergyReportServ) AllergyReportHandler {
	return &allergyReportHandler{serv: serv}
}

func (h *allergyReportHandler) CreateAllergyReport(c *gin.Context) {
	var allergyReportInput api.AllergyReportInput
	if err := c.ShouldBind(&allergyReportInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error brewuu": err.Error()})
		return
	}

	log.Println("allergy report : ", allergyReportInput)

	isExist := h.serv.CheckIsReportExist(allergyReportInput.UserID)
	if isExist {
		c.JSON(http.StatusBadRequest, gin.H{"error brewuu 2": "please wait the confirmation!"})
		return
	}

	allergyReport := models.AllergyReport{
		UserID:    allergyReportInput.UserID,
		Allergies: allergyReportInput.Allergies,
		File:      allergyReportInput.File,
	}

	if err := h.serv.CreateAllergyReport(&allergyReport); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send allergy report", "err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": "Successfully sent allergy report",
		"data":    allergyReport,
	})
}
